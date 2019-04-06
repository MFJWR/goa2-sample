// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// Viron service
//
// Command:
// $ goa gen github.com/tonouchi510/goa2-sample/design

package viron

import (
	"context"

	vironviews "github.com/tonouchi510/goa2-sample/gen/viron/views"
)

// Service is the Viron service interface.
type Service interface {
	// Add viron_authtype
	Authtype(context.Context) (res VironAuthtypeCollection, err error)
	// Add viron_menu
	VironMenu(context.Context) (res *VironMenu, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "Viron"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [2]string{"authtype", "viron_menu"}

// VironAuthtypeCollection is the result type of the Viron service authtype
// method.
type VironAuthtypeCollection []*VironAuthtype

// VironMenu is the result type of the Viron service viron_menu method.
type VironMenu struct {
	Name      string
	Thumbnail *string
	Tags      []string
	Color     *string
	Theme     *string
	Pages     []*VironPage
	Sections  []*VironSection
}

type VironAuthtype struct {
	// type name
	Type string
	// provider name
	Provider string
	// url
	URL string
	// request method to submit this auth
	Method string
}

type VironPage struct {
	ID         string
	Name       string
	Section    string
	Group      *string
	Components []*VironComponent
}

type VironComponent struct {
	Name           string
	Style          string
	Unit           *string
	Actions        []string
	API            *VironAPI
	Pagination     *bool
	Primary        *string
	TableLabels    []string
	Query          []*VironQuery
	AutoRefreshSec *int32
}

type VironAPI struct {
	Method string
	Path   string
}

type VironQuery struct {
	Key  string
	Type string
}

type VironSection struct {
	ID    string
	Label string
}

// NewVironAuthtypeCollection initializes result type VironAuthtypeCollection
// from viewed result type VironAuthtypeCollection.
func NewVironAuthtypeCollection(vres vironviews.VironAuthtypeCollection) VironAuthtypeCollection {
	var res VironAuthtypeCollection
	switch vres.View {
	case "default", "":
		res = newVironAuthtypeCollection(vres.Projected)
	}
	return res
}

// NewViewedVironAuthtypeCollection initializes viewed result type
// VironAuthtypeCollection from result type VironAuthtypeCollection using the
// given view.
func NewViewedVironAuthtypeCollection(res VironAuthtypeCollection, view string) vironviews.VironAuthtypeCollection {
	var vres vironviews.VironAuthtypeCollection
	switch view {
	case "default", "":
		p := newVironAuthtypeCollectionView(res)
		vres = vironviews.VironAuthtypeCollection{p, "default"}
	}
	return vres
}

// NewVironMenu initializes result type VironMenu from viewed result type
// VironMenu.
func NewVironMenu(vres *vironviews.VironMenu) *VironMenu {
	var res *VironMenu
	switch vres.View {
	case "default", "":
		res = newVironMenu(vres.Projected)
	}
	return res
}

// NewViewedVironMenu initializes viewed result type VironMenu from result type
// VironMenu using the given view.
func NewViewedVironMenu(res *VironMenu, view string) *vironviews.VironMenu {
	var vres *vironviews.VironMenu
	switch view {
	case "default", "":
		p := newVironMenuView(res)
		vres = &vironviews.VironMenu{p, "default"}
	}
	return vres
}

// newVironAuthtypeCollection converts projected type VironAuthtypeCollection
// to service type VironAuthtypeCollection.
func newVironAuthtypeCollection(vres vironviews.VironAuthtypeCollectionView) VironAuthtypeCollection {
	res := make(VironAuthtypeCollection, len(vres))
	for i, n := range vres {
		res[i] = newVironAuthtype(n)
	}
	return res
}

// newVironAuthtypeCollectionView projects result type VironAuthtypeCollection
// into projected type VironAuthtypeCollectionView using the "default" view.
func newVironAuthtypeCollectionView(res VironAuthtypeCollection) vironviews.VironAuthtypeCollectionView {
	vres := make(vironviews.VironAuthtypeCollectionView, len(res))
	for i, n := range res {
		vres[i] = newVironAuthtypeView(n)
	}
	return vres
}

// newVironAuthtype converts projected type VironAuthtype to service type
// VironAuthtype.
func newVironAuthtype(vres *vironviews.VironAuthtypeView) *VironAuthtype {
	res := &VironAuthtype{}
	if vres.Type != nil {
		res.Type = *vres.Type
	}
	if vres.Provider != nil {
		res.Provider = *vres.Provider
	}
	if vres.URL != nil {
		res.URL = *vres.URL
	}
	if vres.Method != nil {
		res.Method = *vres.Method
	}
	return res
}

// newVironAuthtypeView projects result type VironAuthtype into projected type
// VironAuthtypeView using the "default" view.
func newVironAuthtypeView(res *VironAuthtype) *vironviews.VironAuthtypeView {
	vres := &vironviews.VironAuthtypeView{
		Type:     &res.Type,
		Provider: &res.Provider,
		URL:      &res.URL,
		Method:   &res.Method,
	}
	return vres
}

// newVironMenu converts projected type VironMenu to service type VironMenu.
func newVironMenu(vres *vironviews.VironMenuView) *VironMenu {
	res := &VironMenu{
		Thumbnail: vres.Thumbnail,
		Color:     vres.Color,
		Theme:     vres.Theme,
	}
	if vres.Name != nil {
		res.Name = *vres.Name
	}
	if vres.Tags != nil {
		res.Tags = make([]string, len(vres.Tags))
		for i, val := range vres.Tags {
			res.Tags[i] = val
		}
	}
	if vres.Pages != nil {
		res.Pages = make([]*VironPage, len(vres.Pages))
		for i, val := range vres.Pages {
			res.Pages[i] = &VironPage{
				ID:      *val.ID,
				Name:    *val.Name,
				Section: *val.Section,
				Group:   val.Group,
			}
			if val.Components != nil {
				res.Pages[i].Components = make([]*VironComponent, len(val.Components))
				for j, val := range val.Components {
					res.Pages[i].Components[j] = &VironComponent{
						Name:           *val.Name,
						Style:          *val.Style,
						Unit:           val.Unit,
						Pagination:     val.Pagination,
						Primary:        val.Primary,
						AutoRefreshSec: val.AutoRefreshSec,
					}
					if val.Actions != nil {
						res.Pages[i].Components[j].Actions = make([]string, len(val.Actions))
						for k, val := range val.Actions {
							res.Pages[i].Components[j].Actions[k] = val
						}
					}
					if val.API != nil {
						res.Pages[i].Components[j].API = transformVironviewsVironAPIViewToVironAPI(val.API)
					}
					if val.TableLabels != nil {
						res.Pages[i].Components[j].TableLabels = make([]string, len(val.TableLabels))
						for k, val := range val.TableLabels {
							res.Pages[i].Components[j].TableLabels[k] = val
						}
					}
					if val.Query != nil {
						res.Pages[i].Components[j].Query = make([]*VironQuery, len(val.Query))
						for k, val := range val.Query {
							res.Pages[i].Components[j].Query[k] = &VironQuery{
								Key:  *val.Key,
								Type: *val.Type,
							}
						}
					}
				}
			}
		}
	}
	return res
}

// newVironMenuView projects result type VironMenu into projected type
// VironMenuView using the "default" view.
func newVironMenuView(res *VironMenu) *vironviews.VironMenuView {
	vres := &vironviews.VironMenuView{
		Name:      &res.Name,
		Thumbnail: res.Thumbnail,
		Color:     res.Color,
		Theme:     res.Theme,
	}
	if res.Tags != nil {
		vres.Tags = make([]string, len(res.Tags))
		for i, val := range res.Tags {
			vres.Tags[i] = val
		}
	}
	if res.Pages != nil {
		vres.Pages = make([]*vironviews.VironPageView, len(res.Pages))
		for i, val := range res.Pages {
			vres.Pages[i] = &vironviews.VironPageView{
				ID:      &val.ID,
				Name:    &val.Name,
				Section: &val.Section,
				Group:   val.Group,
			}
			if val.Components != nil {
				vres.Pages[i].Components = make([]*vironviews.VironComponentView, len(val.Components))
				for j, val := range val.Components {
					vres.Pages[i].Components[j] = &vironviews.VironComponentView{
						Name:           &val.Name,
						Style:          &val.Style,
						Unit:           val.Unit,
						Pagination:     val.Pagination,
						Primary:        val.Primary,
						AutoRefreshSec: val.AutoRefreshSec,
					}
					if val.Actions != nil {
						vres.Pages[i].Components[j].Actions = make([]string, len(val.Actions))
						for k, val := range val.Actions {
							vres.Pages[i].Components[j].Actions[k] = val
						}
					}
					if val.API != nil {
						vres.Pages[i].Components[j].API = transformVironAPIToVironviewsVironAPIView(val.API)
					}
					if val.TableLabels != nil {
						vres.Pages[i].Components[j].TableLabels = make([]string, len(val.TableLabels))
						for k, val := range val.TableLabels {
							vres.Pages[i].Components[j].TableLabels[k] = val
						}
					}
					if val.Query != nil {
						vres.Pages[i].Components[j].Query = make([]*vironviews.VironQueryView, len(val.Query))
						for k, val := range val.Query {
							vres.Pages[i].Components[j].Query[k] = &vironviews.VironQueryView{
								Key:  &val.Key,
								Type: &val.Type,
							}
						}
					}
				}
			}
		}
	}
	return vres
}

// newVironPage converts projected type VironPage to service type VironPage.
func newVironPage(vres *vironviews.VironPageView) *VironPage {
	res := &VironPage{
		Group: vres.Group,
	}
	if vres.Section != nil {
		res.Section = *vres.Section
	}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.Name != nil {
		res.Name = *vres.Name
	}
	if vres.Components != nil {
		res.Components = make([]*VironComponent, len(vres.Components))
		for i, val := range vres.Components {
			res.Components[i] = &VironComponent{
				Name:           *val.Name,
				Style:          *val.Style,
				Unit:           val.Unit,
				Pagination:     val.Pagination,
				Primary:        val.Primary,
				AutoRefreshSec: val.AutoRefreshSec,
			}
			if val.Actions != nil {
				res.Components[i].Actions = make([]string, len(val.Actions))
				for j, val := range val.Actions {
					res.Components[i].Actions[j] = val
				}
			}
			if val.API != nil {
				res.Components[i].API = transformVironviewsVironAPIViewToVironAPI(val.API)
			}
			if val.TableLabels != nil {
				res.Components[i].TableLabels = make([]string, len(val.TableLabels))
				for j, val := range val.TableLabels {
					res.Components[i].TableLabels[j] = val
				}
			}
			if val.Query != nil {
				res.Components[i].Query = make([]*VironQuery, len(val.Query))
				for j, val := range val.Query {
					res.Components[i].Query[j] = &VironQuery{
						Key:  *val.Key,
						Type: *val.Type,
					}
				}
			}
		}
	}
	return res
}

// newVironPageView projects result type VironPage into projected type
// VironPageView using the "default" view.
func newVironPageView(res *VironPage) *vironviews.VironPageView {
	vres := &vironviews.VironPageView{
		ID:      &res.ID,
		Name:    &res.Name,
		Section: &res.Section,
		Group:   res.Group,
	}
	if res.Components != nil {
		vres.Components = make([]*vironviews.VironComponentView, len(res.Components))
		for i, val := range res.Components {
			vres.Components[i] = &vironviews.VironComponentView{
				Name:           &val.Name,
				Style:          &val.Style,
				Unit:           val.Unit,
				Pagination:     val.Pagination,
				Primary:        val.Primary,
				AutoRefreshSec: val.AutoRefreshSec,
			}
			if val.Actions != nil {
				vres.Components[i].Actions = make([]string, len(val.Actions))
				for j, val := range val.Actions {
					vres.Components[i].Actions[j] = val
				}
			}
			if val.API != nil {
				vres.Components[i].API = transformVironAPIToVironviewsVironAPIView(val.API)
			}
			if val.TableLabels != nil {
				vres.Components[i].TableLabels = make([]string, len(val.TableLabels))
				for j, val := range val.TableLabels {
					vres.Components[i].TableLabels[j] = val
				}
			}
			if val.Query != nil {
				vres.Components[i].Query = make([]*vironviews.VironQueryView, len(val.Query))
				for j, val := range val.Query {
					vres.Components[i].Query[j] = &vironviews.VironQueryView{
						Key:  &val.Key,
						Type: &val.Type,
					}
				}
			}
		}
	}
	return vres
}

// newVironComponent converts projected type VironComponent to service type
// VironComponent.
func newVironComponent(vres *vironviews.VironComponentView) *VironComponent {
	res := &VironComponent{
		Unit:           vres.Unit,
		Pagination:     vres.Pagination,
		Primary:        vres.Primary,
		AutoRefreshSec: vres.AutoRefreshSec,
	}
	if vres.Name != nil {
		res.Name = *vres.Name
	}
	if vres.Style != nil {
		res.Style = *vres.Style
	}
	if vres.Actions != nil {
		res.Actions = make([]string, len(vres.Actions))
		for i, val := range vres.Actions {
			res.Actions[i] = val
		}
	}
	if vres.TableLabels != nil {
		res.TableLabels = make([]string, len(vres.TableLabels))
		for i, val := range vres.TableLabels {
			res.TableLabels[i] = val
		}
	}
	if vres.Query != nil {
		res.Query = make([]*VironQuery, len(vres.Query))
		for i, val := range vres.Query {
			res.Query[i] = &VironQuery{
				Key:  *val.Key,
				Type: *val.Type,
			}
		}
	}
	if vres.API != nil {
		res.API = newVironAPI(vres.API)
	}
	return res
}

// newVironComponentView projects result type VironComponent into projected
// type VironComponentView using the "default" view.
func newVironComponentView(res *VironComponent) *vironviews.VironComponentView {
	vres := &vironviews.VironComponentView{
		Name:           &res.Name,
		Style:          &res.Style,
		Unit:           res.Unit,
		Pagination:     res.Pagination,
		Primary:        res.Primary,
		AutoRefreshSec: res.AutoRefreshSec,
	}
	if res.Actions != nil {
		vres.Actions = make([]string, len(res.Actions))
		for i, val := range res.Actions {
			vres.Actions[i] = val
		}
	}
	if res.TableLabels != nil {
		vres.TableLabels = make([]string, len(res.TableLabels))
		for i, val := range res.TableLabels {
			vres.TableLabels[i] = val
		}
	}
	if res.Query != nil {
		vres.Query = make([]*vironviews.VironQueryView, len(res.Query))
		for i, val := range res.Query {
			vres.Query[i] = &vironviews.VironQueryView{
				Key:  &val.Key,
				Type: &val.Type,
			}
		}
	}
	if res.API != nil {
		vres.API = newVironAPIView(res.API)
	}
	return vres
}

// newVironAPI converts projected type VironAPI to service type VironAPI.
func newVironAPI(vres *vironviews.VironAPIView) *VironAPI {
	res := &VironAPI{}
	if vres.Method != nil {
		res.Method = *vres.Method
	}
	if vres.Path != nil {
		res.Path = *vres.Path
	}
	return res
}

// newVironAPIView projects result type VironAPI into projected type
// VironAPIView using the "default" view.
func newVironAPIView(res *VironAPI) *vironviews.VironAPIView {
	vres := &vironviews.VironAPIView{
		Method: &res.Method,
		Path:   &res.Path,
	}
	return vres
}

// newVironQuery converts projected type VironQuery to service type VironQuery.
func newVironQuery(vres *vironviews.VironQueryView) *VironQuery {
	res := &VironQuery{}
	if vres.Key != nil {
		res.Key = *vres.Key
	}
	if vres.Type != nil {
		res.Type = *vres.Type
	}
	return res
}

// newVironQueryView projects result type VironQuery into projected type
// VironQueryView using the "default" view.
func newVironQueryView(res *VironQuery) *vironviews.VironQueryView {
	vres := &vironviews.VironQueryView{
		Key:  &res.Key,
		Type: &res.Type,
	}
	return vres
}

// newVironSection converts projected type VironSection to service type
// VironSection.
func newVironSection(vres *vironviews.VironSectionView) *VironSection {
	res := &VironSection{}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.Label != nil {
		res.Label = *vres.Label
	}
	return res
}

// newVironSectionView projects result type VironSection into projected type
// VironSectionView using the "default" view.
func newVironSectionView(res *VironSection) *vironviews.VironSectionView {
	vres := &vironviews.VironSectionView{
		ID:    &res.ID,
		Label: &res.Label,
	}
	return vres
}

// transformVironviewsVironAPIViewToVironAPI builds a value of type *VironAPI
// from a value of type *vironviews.VironAPIView.
func transformVironviewsVironAPIViewToVironAPI(v *vironviews.VironAPIView) *VironAPI {
	res := &VironAPI{
		Method: *v.Method,
		Path:   *v.Path,
	}

	return res
}

// transformVironAPIToVironviewsVironAPIView builds a value of type
// *vironviews.VironAPIView from a value of type *VironAPI.
func transformVironAPIToVironviewsVironAPIView(v *VironAPI) *vironviews.VironAPIView {
	res := &vironviews.VironAPIView{
		Method: &v.Method,
		Path:   &v.Path,
	}

	return res
}
