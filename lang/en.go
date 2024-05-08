package lang

var en = map[string]string{
	"nousername":                 "Username is required",
	"jsonBreak":                  "JSON is corrupted, try deleting gml.json",
	"username":                   "Username",
	"emailusage":                 "Official/external login account email. Set this when using official/external login, then no need to set username",
	"Downloadusage":              "Version to attempt downloading",
	"Downloadserverusage":        "Download server with main version",
	"Downloadfabricusage":        "Fabric version to attempt downloading",
	"Downloadquiltusage":         "Quilt version to attempt downloading",
	"Downloadpaperusage":         "Paper version to attempt downloading",
	"verlistusage":               "Display all available download versions, for example, release. Use ? to see all optional parameters.",
	"verlistfabricusage":         "Display all available fabric loader versions",
	"verlistquiltusage":          "Display all available quilt loader versions",
	"verlistpaperusage":          "Display all available paper server versions",
	"intusage":                   "Number of goroutines to use when downloading files.",
	"runusage":                   "Version to attempt launching",
	"runpaperusage":              "Paper server to attempt launching",
	"runlistusage":               "Display all available launchable versions",
	"runlistpaperusage":          "Display all available launchable paper servers",
	"ramusage":                   "Allocate memory size (MB) for launching the game",
	"flagusage":                  "Custom launch parameters, such as -XX:+AggressiveOpts -XX:+UseCompressedOops",
	"proxyusage":                 "Set the proxy (http) for downloads",
	"typeusage":                  `Set the download source. Options are vanilla, bmclapi. If not set, it will automatically choose a download source for each file. You can use "bmclapi|vanilla" to load balance and use multiple download sources.`,
	"Independentusage":           "Enable version isolation",
	"testusage":                  "Check file integrity and correctness before launching the game",
	"creditsusage":               "Credits to the projects used",
	"updateusage":                "Check for updates",
	"logusage":                   "Output game logs",
	"yggdrasilusage":             "External login address (authlib-injector)",
	"listusage":                  "View all saved official/external login accounts",
	"removeusage":                "Delete a saved account",
	"javapathusage":              "Set the specified Java path, usually not needed to set this",
	"msusage":                    "Use Microsoft account for login",
	"vusage":                     "View launcher version",
	"langusage":                  "Launcher prompt text language. Defaults to system settings. Options are zh (Chinese), en (English).",
	"bmclapiinfo":                "Uses bmclapi as a mirror download source. Address: https://bmclapidoc.bangbang93.com/",
	"authlib-injectorinfo":       "Uses authlib-injector. Address: https://github.com/yushijinhun/authlib-injector",
	"useproject":                 "Open-source projects used and their licenses can be found at https://github.com/xmdhs/gomclauncher/blob/master/go.mod",
	"checkupdateerr":             "Failed to check for updates",
	"checkupdate":                "Update detected, new version is",
	"nowversion":                 "Current version is",
	"updateinfo":                 "Update details:",
	"nofindLanguage":             "Language not found",
	"authlibdownloadfail":        "Authlib download failed, retrying",
	"authlibcheckerr":            "Authlib verification error, retrying",
	"weberr":                     "Seems to be a network issue, retrying",
	"filecheckerr":               "File verification failed, redownloading",
	"getversionlistfail":         "Failed to get version list, retrying",
	"getfabricversionlistfail":   "Failed to get fabric version list, retrying",
	"getquiltversionlistfail":    "Failed to get quilt version list, retrying",
	"ErrNotSelctProFile":         "Please select a profile, specify with the -username parameter",
	"ErrProFileNoExist":          "Profile does not exist",
	"verifygamejar":              "Verifying game core",
	"downloadgamejar":            "Downloading game core",
	"finish":                     "Completed",
	"verifylibrarie":             "Verifying library files",
	"downloadlibrarie":           "Downloading library files",
	"verifyassets":               "Verifying resource files",
	"downloadassets":             "Downloading resource files",
	"verifynatives":              "Verifying extracted natives library",
	"downloadnatives":            "Downloading and extracting natives library",
	"downloadfail":               "Download failed",
	"download.NoSuch":            "Version not found",
	"download.FileDownLoadFail":  "Too many download failures. Try switching download sources or retrying",
	"msemailnil":                 "Although you don't actually need to enter your email, it is required to mark the account to save the access token. Of course, you can also enter any string that is not an email",
	"msauth.ErrHostname":         "Don't bother with things like forgetting passwords here",
	"auth.ErrCode":               "Attempt to re-login with Microsoft account",
	"auth.ErrProfile":            "It seems you haven't purchased the game or migrated your account yet",
	"msauth.ErrNotInstallChrome": "Please install Chrome, you can download it here https://www.google.cn/intl/zh-CN/chrome/ \n Of course, Chromium is also acceptable",
	"emailnil":                   "Please set the email \n For example, -email example@example.com",
	"auth.NotOk-refresh":         "Please try logging in again",
	"Refresherr":                 "Login failed, possibly due to network issues. You can try again",
	"auth.NotOk":                 "Account name or password is incorrect",
	"namenil":                    "Please create a profile",
	"minecraftlogin":             "Official login",
	"mslogin":                    "Microsoft account login",
	"authlib-injectorlogin":      "External login, API address",
	"email":                      "Email:",
	"name":                       "Username:",
	"nofind":                     "Not found",
	"removeok":                   "Successfully removed",
	"nofindthisversion":          "This version is not installed or there is another issue",
	"flag.os.ErrNotExist":        "Please install the corresponding original version first",
	"launcher.JsonErr":           "JSON error, try deleting the corresponding json file in %v/versions\n",
	"launcher.JsonNorTrue":       "This version's json is incorrect",
	"runlist":                    "Options: ",
	"authlibdownloadfailed":      "Authlib-injector download failed",
	"auth.JsonNotTrue":           "External login address is incorrect",
	"webfail":                    "Perhaps it's a network issue",
	"legacynoexit":               "Resource files do not exist, please enable file verification: %w",
	"tidy":                       "Remove unused files in the assets/objects folder",
	"ErrChild":                   "Underage account, try changing the birthdate or adding to family https://support.xbox.com/zh-cn/help/family-online-safety/child-accounts/add-family-member-on-xbox-one",
}
