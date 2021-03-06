package packages

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// TODO: Replace all home directories with GetConfigDirectory

func FindInstallationYears(softwareLabel SoftwareModule) ([]string, error) {
	var years []string

	files, err := os.ReadDir(app.HomeDirectory + "/Library/Preferences") // gets list of all plist file names
	if err != nil {
		return nil, errors.New("error: could not get files from local library/Preferences")
	}

	// returns all license year numbers found in plist file names from the files variable
	for _, f := range files {
		file := strings.Contains(f.Name(), strings.ToLower(softwareLabel+".license."))
		if file {
			year := regexp.MustCompile("[0-9]+").FindString(f.Name())
			if year != "" {
				years = append(years, year)
			}
		}
	}

	fmt.Println(years)

	return years, nil
}

// setProperties will take in an installation and assign it's properties strings
func (installation *Installation) setProperties() {
	switch installation.SoftwareModule {
	case ModuleVectorworks:
		installation.Properties = []string{
			"net.nemetschek.vectorworks.license." + installation.Year + ".plist",
			"net.nemetschek.vectorworks." + installation.Year + ".plist",
			"net.nemetschek.vectorworks.spotlightimporter.plist",
			"net.nemetschek.vectorworks.plist",
			"net.nemetschek.vectorworksinstaller.helper.plist",
			"net.nemetschek.vectorworksinstaller.plist",
			"net.vectorworks.vectorworks." + installation.Year + ".plist",
		}
	case ModulesVision:
		installation.Properties = []string{
			"com.qtproject.plist",
			"com.vwvision.Vision" + installation.Year + ".plist",
			"com.yourcompany.Vision.plist",
			"net.vectorworks.Vision.plist",
			"net.vectorworks.vision.license." + installation.Year + ".plist",
		}
	}
}

// setUserData well set all user data based on the target packages
func (installation *Installation) setUserData() {
	switch installation.SoftwareModule {
	case ModuleVectorworks:
		installation.Directories = []string{
			app.HomeDirectory + "/Library/Application\\ Support/Vectorworks\\ Cloud\\ Services",
			app.HomeDirectory + "/Library/Application\\ Support/Vectorworks/" + installation.Year,
			app.HomeDirectory + "/Library/Application\\ Support/vectorworks-installer-wrapper",
		}
	case ModulesVision:
		installation.Directories = []string{
			app.HomeDirectory + "/Library/Application\\ Support/Vision/" + installation.Year,
			app.HomeDirectory + "/Library/Application\\ Support/VisionUpdater",
			"/Library/Frameworks/QtConcurrent.framework",
			"/Library/Frameworks/QtCore.framework",
			"/Library/Frameworks/QtDBus.framework",
			"/Library/Frameworks/QtGui.framework",
			"/Library/Frameworks/QtNetwork.framework",
			"/Library/Frameworks/QtOpenGL.framework",
			"/Library/Frameworks/QtPlugins",
			"/Library/Frameworks/QtPositioning.framework",
			"/Library/Frameworks/QtPrintSupport.framework",
			"/Library/Frameworks/QtQml.framework",
			"/Library/Frameworks/QtQuick.framework",
			"/Library/Frameworks/QtWebChannel.framework",
			"/Library/Frameworks/QtWebEngine.framework",
			"/Library/Frameworks/QtWebEngineCore.framework",
			"/Library/Frameworks/QtWebEngineWidgets.framework",
			"/Library/Frameworks/QtWidgets.framework",
			"/Library/Frameworks/QtXml.framework",
			"/Library/Frameworks/rpath_manipulator.sh",
			"/Library/Frameworks/setup_qt_frameworks.sh",
		}
	}
}

// setRMCache sets the system path for the resource manager cache directory
func (installation *Installation) setRMCache() {
	installation.RMCache = app.HomeDirectory + "/Library/Application\\ Support/Vectorworks\\ RMCache/rm" + installation.Year
}

// setLogFiles sets all the log files paths for the target packages
func (installation *Installation) setLogFileSent() {
	installation.LogFileSent = app.HomeDirectory + "/Library/Application\\ Support/Vectorworks/" + installation.Year + "/VW User Log Sent.txt"
}

// setLogFiles sets all the log files paths for the target packages
func (installation *Installation) setLogFile() {
	installation.LogFile = app.HomeDirectory + "/Library/Application\\ Support/Vectorworks/" + installation.Year + "/VW User Log.txt"
}

//
// CLEANING SECTION
//

func (installation *Installation) cleanRMCache() {
	// check that a directory exists
	if _, err := os.Stat(installation.RMCache); os.IsNotExist(err) {
		app.SoftwareOutputString = append(app.SoftwareOutputString, "No Resource Manager Cache found.")
		_ = errors.New("RMCache directory does not exist")
	}

	// Attempt to remove the directory
	err := os.RemoveAll(installation.RMCache)
	if err != nil {
		app.SoftwareOutputString = append(app.SoftwareOutputString, "error: could not delete the directory at: "+installation.RMCache)
		_ = errors.New("error: could not delete the directory: " + installation.RMCache)
	}
	app.SoftwareOutputString = append(app.SoftwareOutputString, fmt.Sprintf("%s - deleted the directory: %s", installation.SoftwareModule, installation.RMCache))
}

func (installation *Installation) cleanUserData() {
	for _, directory := range installation.Directories {

		// check that a directory exists
		if _, err := os.Stat(directory); os.IsNotExist(err) {
			app.SoftwareOutputString = append(app.SoftwareOutputString, "No directory found at: "+directory)
		}

		app.SoftwareOutputString = append(app.SoftwareOutputString, fmt.Sprintf("Cleaning %s...", directory))

		err := os.RemoveAll(directory)
		if err != nil {
			app.SoftwareOutputString = append(app.SoftwareOutputString, "error: could not delete the directory at: "+directory)
			_ = errors.New("error: could not delete the directory: " + directory)
		}
		app.SoftwareOutputString = append(app.SoftwareOutputString, fmt.Sprintf("%s - deleted the directory: %s", installation.SoftwareModule, directory))
	}
}

func (installation *Installation) cleanUserSettings() {
	for _, property := range installation.Properties {
		// check that a property exists
		if _, err := os.Stat(app.HomeDirectory + "/Library/Preferences/" + property); os.IsNotExist(err) {
			app.SoftwareOutputString = append(app.SoftwareOutputString, "No property found at: "+property)
		}

		app.SoftwareOutputString = append(app.SoftwareOutputString, fmt.Sprintf("Cleaning %s...", property))

		err := os.Remove(app.HomeDirectory + "/Library/Preferences/" + property)
		if err != nil {
			app.SoftwareOutputString = append(app.SoftwareOutputString, "error: could not delete the property at: "+property)
			_ = errors.New("error: could not delete the property: " + property)
		}
		app.SoftwareOutputString = append(app.SoftwareOutputString, fmt.Sprintf("%s - deleted the property: %s", installation.SoftwareModule, property))
	}
}

func (installation *Installation) cleanInstallers() error {
	//for _, installer := range installation.Installers {
	//	err := os.Remove(installer)
	//	if err != nil {
	//		return errors.New("error: could not delete the installer: " + installer)
	//	}
	//}

	// Not yet implemented
	// TODO: implement this function

	return nil
}

func (installation *Installation) cleanAll() {
	installation.cleanRMCache()
	installation.cleanUserData()
	installation.cleanUserSettings()
	_ = installation.cleanInstallers()
}

func (installation Installation) Clean() error {
	app.SoftwareOutputString = append(app.SoftwareOutputString, fmt.Sprintf("%s - cleaning...", installation.SoftwareModule))

	if installation.CleanOptions.RemoveAllData {
		app.SoftwareOutputString = append(app.SoftwareOutputString, "Cleaning all data...")
		installation.cleanAll()
		return nil
	}

	if installation.CleanOptions.RemoveRMC {
		app.SoftwareOutputString = append(app.SoftwareOutputString, "Cleaning Resource Manager Cache...")
		installation.cleanRMCache()
	}

	if installation.CleanOptions.RemoveUserData {
		app.SoftwareOutputString = append(app.SoftwareOutputString, "Cleaning user data...")
		installation.cleanUserData()
	}

	if installation.CleanOptions.RemoveUserSettings {
		app.SoftwareOutputString = append(app.SoftwareOutputString, "Cleaning user settings...")
		installation.cleanUserSettings()
	}

	if installation.CleanOptions.RemoveInstallerSettings {
		app.SoftwareOutputString = append(app.SoftwareOutputString, "Cleaning installers...")
		err := installation.cleanInstallers()
		if err != nil {
			return err
		}
	}

	return nil
}
