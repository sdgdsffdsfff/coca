package tbs

import (
	"github.com/phodal/coca/core/models"
	"strings"
)

type TbsApp struct {
}

var (
	DuplicatedAssertionLimitLength = 5
)

func NewTbsApp() *TbsApp {
	return &TbsApp{}
}

type TestBadSmell struct {
	FileName    string
	Type        string
	Description string
	Line        int
}

func (a TbsApp) AnalysisPath(deps []models.JClassNode, identifiersMap map[string]models.JIdentifier) []TestBadSmell {
	var results []TestBadSmell = nil

	var callMethodMap = make(map[string]models.JMethod)
	for _, clz := range deps {
		for _, method := range clz.Methods {
			callMethodMap[clz.Package+"."+clz.Class+"."+method.Name] = method
		}
	}

	for _, clz := range deps {
		// TODO refactoring identify & annotation
		for _, method := range clz.Methods {
			if !isTest(method) {
				continue
			}

			currentMethodCalls := updateMethodCalls(method, clz, callMethodMap)

			var testType = ""
			for _, annotation := range method.Annotations {
				checkIgnoreTest(clz.Path, annotation, &results, &testType)
				checkEmptyTest(clz.Path, annotation, &results, method, &testType)
			}

			var methodCallMap = make(map[string][]models.JMethodCall)
			var hasAssert = false
			for index, methodCall := range currentMethodCalls {
				if methodCall.MethodName == "" {
					continue
				}
				methodCallMap[getMethodCallFullPath(methodCall)] = append(methodCallMap[getMethodCallFullPath(methodCall)], methodCall)

				checkRedundantPrintTest(clz.Path, methodCall, &results, &testType)
				checkSleepyTest(clz.Path, methodCall, method, &results, &testType)
				checkRedundantAssertionTest(clz.Path, methodCall, method, &results, &testType)

				if hasAssertion(methodCall.MethodName) {
					hasAssert = true
				}

				if index == len(currentMethodCalls)-1 {
					if !hasAssert {
						checkUnknownTest(clz, method, &results, &testType)
					}
				}
			}

			checkDuplicateAssertTest(clz, &results, methodCallMap, method, &testType)
		}
	}

	return results
}

func updateMethodCalls(method models.JMethod, clz models.JClassNode, callMethodMap map[string]models.JMethod) []models.JMethodCall {
	currentMethodCalls := method.MethodCalls
	for _, methodCall := range currentMethodCalls {
		if methodCall.Class == clz.Class {
			jMethod := callMethodMap[getMethodCallFullPath(methodCall)]
			if jMethod.Name != "" {
				currentMethodCalls = append(currentMethodCalls, jMethod.MethodCalls...)
			}
		}
	}
	return currentMethodCalls
}

func checkRedundantAssertionTest(path string, call models.JMethodCall, method models.JMethod, results *[]TestBadSmell, testType *string) {
	TWO_PARAMETERS := 2
	if len(call.Parameters) == TWO_PARAMETERS {
		if call.Parameters[0] == call.Parameters[1] {
			*testType = "RedundantAssertionTest"
			tbs := *&TestBadSmell{
				FileName:    path,
				Type:        *testType,
				Description: "",
				Line:        method.StartLine,
			}

			*results = append(*results, tbs)
		}
	}
}

func hasAssertion(methodName string) bool {
	methodName = strings.ToLower(methodName)
	assertionList := []string{
		"assert",
		"should",
		"check",    // ArchUnit,
		"maynotbe", // ArchUnit,
		"is",       // RestAssured,
		"spec",     // RestAssured,
		"verify",   // Mockito,
	}

	for _, assertion := range assertionList {
		if strings.Contains(methodName, assertion) {
			return true
		}
	}

	return false
}

func isTest(method models.JMethod) bool {
	var isTest = false
	for _, annotation := range method.Annotations {
		if annotation.QualifiedName == "Test" || annotation.QualifiedName == "Ignore" {
			isTest = true
		}
	}
	return isTest
}

func checkUnknownTest(clz models.JClassNode, method models.JMethod, results *[]TestBadSmell, testType *string) {
	*testType = "UnknownTest"
	tbs := *&TestBadSmell{
		FileName:    clz.Path,
		Type:        *testType,
		Description: "",
		Line:        method.StartLine,
	}

	*results = append(*results, tbs)
}

func checkDuplicateAssertTest(clz models.JClassNode, results *[]TestBadSmell, methodCallMap map[string][]models.JMethodCall, method models.JMethod, testType *string) {
	var isDuplicateAssert = false
	for _, methodCall := range methodCallMap {
		if len(methodCall) >= DuplicatedAssertionLimitLength {
			methodName := methodCall[len(methodCall)-1].MethodName
			if hasAssertion(methodName) {
				isDuplicateAssert = true
			}
		}
	}

	if isDuplicateAssert {
		*testType = "DuplicateAssertTest"
		tbs := *&TestBadSmell{
			FileName:    clz.Path,
			Type:        *testType,
			Description: "",
			Line:        method.StartLine,
		}

		*results = append(*results, tbs)
	}
}

func getMethodCallFullPath(methodCall models.JMethodCall) string {
	return methodCall.Package + "." + methodCall.Class + "." + methodCall.MethodName
}

func checkSleepyTest(path string, method models.JMethodCall, jMethod models.JMethod, results *[]TestBadSmell, testType *string) {
	if method.MethodName == "sleep" && method.Class == "Thread" {
		*testType = "SleepyTest"
		tbs := *&TestBadSmell{
			FileName:    path,
			Type:        *testType,
			Description: "",
			Line:        method.StartLine,
		}

		*results = append(*results, tbs)
	}
}

func checkRedundantPrintTest(path string, mCall models.JMethodCall, results *[]TestBadSmell, testType *string) {
	if mCall.Class == "System.out" && (mCall.MethodName == "println" || mCall.MethodName == "printf" || mCall.MethodName == "print") {
		*testType = "RedundantPrintTest"
		tbs := *&TestBadSmell{
			FileName:    path,
			Type:        *testType,
			Description: "",
			Line:        mCall.StartLine,
		}

		*results = append(*results, tbs)
	}
}

func checkEmptyTest(path string, annotation models.Annotation, results *[]TestBadSmell, method models.JMethod, testType *string) {
	if annotation.QualifiedName == "Test" {
		if len(method.MethodCalls) <= 1 {
			*testType = "EmptyTest"
			tbs := *&TestBadSmell{
				FileName:    path,
				Type:        *testType,
				Description: "",
				Line:        method.StartLine,
			}

			*results = append(*results, tbs)
		}
	}
}

func checkIgnoreTest(clzPath string, annotation models.Annotation, results *[]TestBadSmell, testType *string) {
	if annotation.QualifiedName == "Ignore" {
		*testType = "IgnoreTest"
		tbs := *&TestBadSmell{
			FileName:    clzPath,
			Type:        *testType,
			Description: "",
			Line:        0,
		}

		*results = append(*results, tbs)
	}
}
