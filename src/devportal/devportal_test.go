package devportal

import (
	"testing"
)

func TestGetMetadata(t *testing.T) {
	string, err := getMetadata("https://services.odata.org/V2/Northwind/Northwind.svc/$metadata", "none")
	if err != nil {
		t.Errorf("failed with error: %s", err)
	}
	t.Log(string)
}

func TestMetadatatoOpenAPI(t *testing.T) {
	metaData, err := getMetadata("https://services.odata.org/V2/Northwind/Northwind.svc/$metadata", "none")
	if err != nil {
		t.Errorf("failed with error: %s", err)
	}
	_, err = convertMetaDatatoOpenAPI(metaData)
	if err != nil {
		t.Errorf("failed with error: %s", err)
	}
}

func TestCreateDevPortalProject(t *testing.T) {
	err := createDevPortalProject("test4trace", "description4trace")
	if err != nil {
		t.Errorf("failed with error: %s", err)
	}

	// Calling this twice should work
	err = createDevPortalProject("test4trace", "description4trace")
	if err != nil {
		t.Errorf("failed with error: %s", err)
	}
}
