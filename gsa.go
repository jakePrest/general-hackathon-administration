package main

import (
    "github.com/montanaflynn/stats"
    "net/http"
    "github.com/gorilla/mux"
    "log"
  //  "strconv"
//    "fmt"
    "html/template"
    "encoding/json"
)

type Data1 struct {
/*	DataCenterID string `json:"Data Center ID"`
	AgencyDataCenterID string `json:"Agency Data Center ID"`
	LegacyDataCenterID string `json:"Legacy Data Center ID"`
	RecordStatus string `json:"Record Status"`
	RecordValidity string `json:"Record Validity"`
	AgencyAbbreviation string `json:"Agency Abbreviation"`
	Component string `json:"Component"`
	OfficeName string `json:"Office Name"`
	DataCenterName string `json:"Data Center Name"`
	PublishedName string `json:"Published Name"`
	CoreClassification string `json:"Core Classification"`
	NonCoreSubCategory string `json:"Non-core Sub-Category"`
	City string `json:"City"`
	State string `json:"State"`
	Province string `json:"Province"`
	ZipCode string `json:"Zip Code"`
	Country string `json:"Country"`
	OwnershipType string `json:"Ownership Type"`
	ColocationProvider string `json:"Colocation Provider"`
	ColocationProviderType string `json:"Colocation Provider Type"`
	GovernmentWideProvider string `json:"Government-Wide Provider"`
	ProvidingServices string `json:"Providing Services"`
	PercentServicesPaidByOtherAgencies string `json:"Percent Services Paid by Other Agencies"`
	OtherAgenciesServiced string `json:"Other Agencies Serviced"`
	DataCenterTier string `json:"Data Center Tier"`
	GrossFloorArea string `json:"Gross Floor Area"`
	TotalCustomerFloorArea string `json:"Total Customer Floor Area"`
	AnnualCostPerSqFt string `json:"Annual Cost per Sq Ft"`
	Fte string `json:"FTE"`
	FteCost string `json:"FTE Cost"`
	ElectricityIncludedInCost string `json:"Electricity Included in Cost"`
	ElectricityIsMetered string `json:"Electricity is Metered"`
	TotalPowerCapacity string `json:"Total Power Capacity"`
	TotalItPowerCapacity string `json:"Total IT Power Capacity"`
	AvgElectricityUsage string `json:"Avg Electricity Usage"`
	AvgItElectricityUsage string `json:"Avg IT Electricity Usage"`
	CostPerKwh string `json:"Cost per KWH"`
	RackCount string `json:"Rack Count"`
	TotalIbmMainframes string `json:"Total IBM Mainframes"`
	TotalOtherMainframes string `json:"Total Other Mainframes"`*/
	TotalWindowsServers float64 `json:"Total Windows Servers"`
	/* string `json:"Total Unix Servers"`
	TotalLinuxServers string `json:"Total Linux Servers"`
	TotalHpcClusterNodes string `json:"Total HPC Cluster Nodes"`
	OtherServers string `json:"Other Servers"`
	TotalVirtualHosts string `json:"Total Virtual Hosts"`
	TotalVirtualOs string `json:"Total Virtual OS"`
	TotalStorage string `json:"Total Storage"`
	UsedStorage string `json:"Used Storage"`
	ClosingStage string `json:"Closing Stage"`
	ClosingFiscalYear string `json:"Closing Fiscal Year"`
	ClosingQuarter string `json:"Closing Quarter"`
	RealPropertyDisposition string `json:"Real Property Disposition"`
	RealPropertyDispositionDate string `json:"Real Property Disposition Date"`
	TotalFloorAreaEliminatedOrRepurposed string `json:"Total Floor Area Eliminated or Repurposed"`
	TotalDecommissionedPhysicalServers string `json:"Total Decommissioned Physical Servers"`
	TotalServersMovedToOtherDataCenter string `json:"Total Servers Moved to Other Data Center"`
	OverallFteReduction string `json:"Overall FTE Reduction"` */
}

func Index(w http.ResponseWriter, r *http.Request) {
   t := template.New("index.html")
   t, _ = t.ParseFiles("./index.html")  // Parse template file.
   t.Execute(w, nil)
}

/* ProcessData processes Data from JSON request */
func ProcessData(w http.ResponseWriter, r *http.Request) {
   // Compute Cost of Electricity
   // Compute Gross Space
   // Storage Utilization = Used/Total
   // Storage Utilization Mean
   // FTE * Federal Min Wage
  decoder := json.NewDecoder(r.Body)
  ds := make([]Data1,0)
  err := decoder.Decode(&ds)
  if err != nil {
      log.Println("Processing Error:", err)
  }

  data := make([]float64, 0)
  for i := 0; i < len(ds); i++ {
    data = append(data, ds[i].TotalWindowsServers)
  }
var d stats.Float64Data = data

totalwinsserver, _ := d.Mean()
log.Println(totalwinsserver)

}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/", Index)
    r.HandleFunc("/api/process", ProcessData)
    err := http.ListenAndServe(":7000", r)
    if err != nil {
      log.Fatal(err)
    }
}
