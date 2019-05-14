package devportal

import (
	"os"
	"errors"
	"net/http"
	"io/ioutil"
	"fmt"
	"strings"
	"okta"
	"apiproxy"
	"github.com/apex/log"
)

const devPortalProjectsURL = "https://api-developer.niketech.com/prod/v1/projects"
const devPortalUser = "nike.sapcp.entcat"

func Handler(c chan apiproxy.APIProxy) {
	for {
		api := <- c
		ctx := log.WithFields(log.Fields{
			"api": api.Name,
			"tenant": api.Tenant,
		})
		err := createDevPortalProject(api.Name, api.Description)
		ctx.Trace("DevPortal Creation").Stop(&err)
	}
}

func createDevPortalProject(name, description string) error {
	token, err:= okta.GenerateToken(devPortalUser, os.Getenv("DEVPORTAL_SECRET"), true)
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to Create Dev Portal Project: %s", err))
	}
	client := &http.Client{}

	req, err := http.NewRequest("POST", devPortalProjectsURL, strings.NewReader(createProjectPayload(name, description)))
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", "Bearer "+token)
	resp, err := client.Do(req)
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to Create Dev Portal Project: %s", err))
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if (resp.StatusCode < 200 || resp.StatusCode > 299) && !(strings.Contains(string(bodyBytes), "it already exists")) {
		return errors.New(fmt.Sprintf("Failed to Create Dev Portal Project, returned status code: %s", resp.Status))
	}

	return nil
}

func createProjectPayload(name, description string) string {
	return fmt.Sprintf(
`
{
   "department":"Niketech",
   "domains":[
      "Analytics"
   ],
   "projectName":"%s",
   "issuesUrl":"piyush.uthaman@nike.com",
   "ownerId":"nike.sapcp.entcat",
   "isNativeProject":"yes",
   "latestVersionIndex":1,
   "latestVersion":"1.0",
   "isPublished":"yes",
   "ownerType":"user",
   "maintainerEmail":"piyush.uthaman@nike.com",
   "description":"%s",
   "searchName":"%s",
   "projectType":"api",
   "versions":{
      "1.0":{
	 "defaultEnvironment":"prod",
	 "environments":{
	    "dev":{
	       "name":"dev",
	       "hostUrl":"https://nikescpdev.apimanagement.us2.hana.ondemand.com",
	       "healthCheckConfig":{
		  "type":"none"
	       }
	    },
	    "qa":{
	       "name":"qa",
	       "hostUrl":"https://nikescpqa.apimanagement.us2.hana.ondemand.com",
	       "healthCheckConfig":{
		  "type":"none"
	       }
	    },
	    "prod":{
	       "name":"prod",
	       "hostUrl":"https://nikescpprod.apimanagement.us2.hana.ondemand.com",
	       "healthCheckConfig":{
		  "type":"none"
	       }
	    }
	 },
	 "versionName":"1.0",
	 "slas":{

	 },
	 "hasExternalReadme":false,
	 "hasExternalApi":false,
	 "status":"release",
	 "apiDocumentUrl":"https://github.nike.com/scp/APIM-Backup/blob/master/qa/%s/spec.json"
      }
   }
}`, name, description, name, name)
}
