package cc.arduino.packages.contributions;

import org.junit.Test;
import processing.app.Platform;

import static org.junit.Assert.assertFalse;
import static org.junit.Assert.assertTrue;

public class HostDependentDownloadableContribution{

  public void macOsXPositiveTest() {
    HostDependentDownloadableContributionStub contribution = new HostDependentDownloadableContributionStub() {
      @Override
      public String getHost() {
        return "x86_64-apple-darwin";
      }
    };

    Platform platform = new Platform() {
      @Override
      public String getOsName() {
        return "Mac OS X";
      }

      @Override
      public String getOsArch() {
        return "x86_64";
      }
    };
  }
}