/*
 * Strava API v3
 *
 * Strava API
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type PolylineMap struct {

	// The identifier of the map
	Id string `json:"id,omitempty"`

	// The polyline of the map
	Polyline string `json:"polyline,omitempty"`

	// The summary polyline of the map
	SummaryPolyline string `json:"summary_polyline,omitempty"`
}
