package catalog

import (
	"fmt"
	"testing"

	installplanv1alpha1 "github.com/coreos-inc/alm/apis/installplan/v1alpha1"

	"github.com/coreos/go-semver/semver"
)

// memCatalog is a rough mock catalog that holds apps and declarations in a map
type memCatalog struct {
	versions     map[string]semver.Versions
	declarations map[string]map[string]installplanv1alpha1.InstallPlan
}

func latest(verlist semver.Versions) (*semver.Version, bool) {
	if verlist.Len() < 1 {
		return nil, false
	}
	semver.Sort(verlist)
	return verlist[verlist.Len()-1], true
}

func (cat *memCatalog) FetchLatestVersion(apptype string) (*semver.Version, error) {
	versions, ok := cat.versions[apptype]
	if !ok {
		return nil, fmt.Errorf("unknown apptype: %s", apptype)
	}
	ver, ok := latest(versions)
	if !ok {
		return nil, fmt.Errorf("cannot find valid version for apptype %s", apptype)
	}
	return ver, nil
}

func (cat *memCatalog) FetchInstallPlanForAppVersion(apptype, version string) (*installplanv1alpha1.InstallPlan, error) {
	appversions, ok := cat.declarations[apptype]
	if !ok {
		return nil, fmt.Errorf("unknown apptype: %s", apptype)
	}
	decl, ok := appversions[version]
	if !ok {
		return nil, fmt.Errorf("unknown version %s for app: %s", version, apptype)
	}
	return &decl, nil
}

func (cat *memCatalog) ResolveDependencies(decl *installplanv1alpha1.InstallPlan) error {
	// you don't get no dependencies!
	return nil
}

func TestCatalog(t *testing.T) {
	return
}
