package restConsume

type AvailRequest struct {
	TravelPreferences struct {
		TPAExtensions struct {
			DataSources struct {
				ATPCO string
				LCC   string
				NDC   string
			}
			NumTrips struct {
			}
		}
	}
	TravelerInfoSummary struct {
		AirTravelerAvail []struct {
			PassengerTypeQuantity []struct {
				Code     string
				Quantity int
			}
		}
		SeatsRequested []int
	}
	Version string
}

type OTAAirLowFareSearchRQ struct {
	Version string
}

type OriginDestinationInformation []struct {
	DepartureDateTime   string
	DestinationLocation struct {
		LocationCode string
	}
	OriginLocation struct {
		LocationCode string
	}
	RPH string
}

type POS struct {
	Source []struct {
		PseudoCityCode string
		RequestorID    struct {
			CompanyName struct {
				Code string
			}
			ID   string
			Type string
		}
	}
}

type TPAExtensions struct {
	IntelliSellTransaction struct {
		RequestType struct {
			Name string
		}
	}
}

type TravelPreferences struct {
	TPAExtensions struct {
		DataSources struct {
			ATPCO string
			LCC   string
			NDC   string
		}
		NumTrips struct {
		}
	}
}

type TravelerInfoSummary struct {
	AirTravelerAvail []struct {
		PassengerTypeQuantity []struct {
			Code     string
			Quantity int
		}
	}
	SeatsRequested []int
}

func (req *OriginDestinationInformation) SetDepartureDate(departureDate string) OriginDestinationInformation {

	c := OriginDestinationInformation{

		"DepartureDateTime": departureDate}
	return c

}

func (*POS) setPos(pos string) POS {
	p := POS{
		"source": {
			"PseudoCityCode": pos,
			"RequestorID": {
				"CompanyName": {
					"Code": "TN"},
				"ID":   "1",
				"Type": "1"},
		},
	}
	return p
}

func (*TPAExtensions) tpaextensions() {

	tp = TPAExtensions{
		"IntelliSellTransaction": {
			"RequestType": {
				"Name": "200ITINS",
			},
		},
	}
	return tp
}
