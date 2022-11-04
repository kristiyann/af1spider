package models

import "time"

type NikeProductResponse struct {
	Pages struct {
		Prev           string `json:"prev"`
		Next           string `json:"next"`
		TotalPages     int    `json:"totalPages"`
		TotalResources int    `json:"totalResources"`
	} `json:"pages"`
	Objects []struct {
		ID               string    `json:"id"`
		ChannelID        string    `json:"channelId"`
		ChannelName      string    `json:"channelName"`
		Marketplace      string    `json:"marketplace"`
		Language         string    `json:"language"`
		LastFetchTime    time.Time `json:"lastFetchTime"`
		PublishedContent struct {
			Preview            bool          `json:"preview"`
			ExternalReferences []interface{} `json:"externalReferences"`
			Marketplace        string        `json:"marketplace"`
			CollectionGroupID  string        `json:"collectionGroupId"`
			CreatedDateTime    time.Time     `json:"createdDateTime"`
			Language           string        `json:"language"`
			ViewStartDate      time.Time     `json:"viewStartDate"`
			Type               string        `json:"type"`
			Version            string        `json:"version"`
			Analytics          struct {
				HashKey string `json:"hashKey"`
			} `json:"analytics"`
			Nodes []struct {
				Analytics struct {
					HashKey string `json:"hashKey"`
				} `json:"analytics"`
				Nodes []struct {
					Analytics struct {
						HashKey string `json:"hashKey"`
					} `json:"analytics"`
					SubType    string `json:"subType"`
					ID         string `json:"id"`
					Type       string `json:"type"`
					Version    string `json:"version"`
					Properties struct {
						PortraitID   string `json:"portraitId"`
						SquarishURL  string `json:"squarishURL"`
						LandscapeID  string `json:"landscapeId"`
						AltText      string `json:"altText"`
						PortraitURL  string `json:"portraitURL"`
						LandscapeURL string `json:"landscapeURL"`
						Title        string `json:"title"`
						Squarish     struct {
							View        string  `json:"view"`
							AspectRatio float64 `json:"aspectRatio"`
							ID          string  `json:"id"`
							Type        string  `json:"type"`
							URL         string  `json:"url"`
						} `json:"squarish"`
						Portrait struct {
							View        string  `json:"view"`
							AspectRatio float64 `json:"aspectRatio"`
							ID          string  `json:"id"`
							Type        string  `json:"type"`
							URL         string  `json:"url"`
						} `json:"portrait"`
						SeoName    string `json:"seoName"`
						SquarishID string `json:"squarishId"`
						Subtitle   string `json:"subtitle"`
						ColorTheme string `json:"colorTheme"`
					} `json:"properties,omitempty"`
					Properties0 struct {
						StartImageURL string `json:"startImageURL"`
						VideoURL      string `json:"videoURL"`
						Loop          bool   `json:"loop"`
						ProviderID    string `json:"providerId"`
						Subtitle      string `json:"subtitle"`
						ColorTheme    string `json:"colorTheme"`
						VideoID       string `json:"videoId"`
						AutoPlay      bool   `json:"autoPlay"`
						Title         string `json:"title"`
					} `json:"properties,omitempty"`
				} `json:"nodes,omitempty"`
				SubType    string `json:"subType"`
				ID         string `json:"id"`
				Type       string `json:"type"`
				Version    string `json:"version"`
				Properties struct {
					ContainerType string `json:"containerType"`
					Loop          bool   `json:"loop"`
					Subtitle      string `json:"subtitle"`
					ColorTheme    string `json:"colorTheme"`
					AutoPlay      bool   `json:"autoPlay"`
					Title         string `json:"title"`
					Body          string `json:"body"`
					Speed         int    `json:"speed"`
				} `json:"properties,omitempty"`
				Properties0 struct {
					CopyID        string        `json:"copyId"`
					Product       []interface{} `json:"product"`
					RichTextLinks []interface{} `json:"richTextLinks"`
					Custom        struct {
					} `json:"custom"`
					Subtitle string `json:"subtitle"`
					Style    struct {
						DefaultStyle struct {
							Margin struct {
								Top string `json:"top"`
							} `json:"margin"`
						} `json:"defaultStyle"`
						ModifiedDate   time.Time `json:"modifiedDate"`
						ExposeTemplate bool      `json:"exposeTemplate"`
						Properties     struct {
							Actions struct {
							} `json:"actions"`
						} `json:"properties"`
						ResourceType string `json:"resourceType"`
					} `json:"style"`
					JSONBody struct {
						Type    string `json:"type"`
						Content []struct {
							Type    string `json:"type"`
							Content []struct {
								Text string `json:"text"`
								Type string `json:"type"`
							} `json:"content"`
						} `json:"content"`
					} `json:"jsonBody"`
					Body    string        `json:"body"`
					Title   string        `json:"title"`
					Actions []interface{} `json:"actions"`
				} `json:"properties,omitempty"`
				Properties1 struct {
					SquarishURL  string `json:"squarishURL"`
					AltText      string `json:"altText"`
					PortraitURL  string `json:"portraitURL"`
					LandscapeURL string `json:"landscapeURL"`
					Title        string `json:"title"`
					Portrait     struct {
						View string `json:"view"`
						ID   string `json:"id"`
						Type string `json:"type"`
						URL  string `json:"url"`
					} `json:"portrait"`
					CopyID        string        `json:"copyId"`
					RichTextLinks []interface{} `json:"richTextLinks"`
					Subtitle      string        `json:"subtitle"`
					ColorTheme    string        `json:"colorTheme"`
					Style         struct {
						DefaultStyle struct {
							Margin struct {
								Top string `json:"top"`
							} `json:"margin"`
						} `json:"defaultStyle"`
						ModifiedDate   time.Time `json:"modifiedDate"`
						ExposeTemplate bool      `json:"exposeTemplate"`
						ResourceType   string    `json:"resourceType"`
					} `json:"style"`
					Actions   []interface{} `json:"actions"`
					Landscape struct {
						View string `json:"view"`
						ID   string `json:"id"`
						Type string `json:"type"`
						URL  string `json:"url"`
					} `json:"landscape"`
				} `json:"properties,omitempty"`
				Properties2 struct {
					CopyID        string        `json:"copyId"`
					Product       []interface{} `json:"product"`
					RichTextLinks []interface{} `json:"richTextLinks"`
					Custom        struct {
					} `json:"custom"`
					Subtitle string `json:"subtitle"`
					Style    struct {
						DefaultStyle struct {
							Margin struct {
								Top string `json:"top"`
							} `json:"margin"`
						} `json:"defaultStyle"`
						ModifiedDate   time.Time `json:"modifiedDate"`
						ExposeTemplate bool      `json:"exposeTemplate"`
						Properties     struct {
							Actions struct {
							} `json:"actions"`
						} `json:"properties"`
						ResourceType string `json:"resourceType"`
					} `json:"style"`
					JSONBody struct {
						Type    string `json:"type"`
						Content []struct {
							Type    string `json:"type"`
							Content []struct {
								Text string `json:"text"`
								Type string `json:"type"`
							} `json:"content"`
						} `json:"content"`
					} `json:"jsonBody"`
					Body    string        `json:"body"`
					Title   string        `json:"title"`
					Actions []interface{} `json:"actions"`
				} `json:"properties,omitempty"`
				Properties3 struct {
					SquarishURL  string `json:"squarishURL"`
					AltText      string `json:"altText"`
					PortraitURL  string `json:"portraitURL"`
					LandscapeURL string `json:"landscapeURL"`
					Title        string `json:"title"`
					Portrait     struct {
						View string `json:"view"`
						ID   string `json:"id"`
						Type string `json:"type"`
						URL  string `json:"url"`
					} `json:"portrait"`
					CopyID        string        `json:"copyId"`
					RichTextLinks []interface{} `json:"richTextLinks"`
					Subtitle      string        `json:"subtitle"`
					ColorTheme    string        `json:"colorTheme"`
					Style         struct {
						DefaultStyle struct {
							Margin struct {
								Top string `json:"top"`
							} `json:"margin"`
						} `json:"defaultStyle"`
						ModifiedDate   time.Time `json:"modifiedDate"`
						ExposeTemplate bool      `json:"exposeTemplate"`
						ResourceType   string    `json:"resourceType"`
					} `json:"style"`
					Actions   []interface{} `json:"actions"`
					Landscape struct {
						View string `json:"view"`
						ID   string `json:"id"`
						Type string `json:"type"`
						URL  string `json:"url"`
					} `json:"landscape"`
				} `json:"properties,omitempty"`
				Properties4 struct {
					CopyID        string        `json:"copyId"`
					Product       []interface{} `json:"product"`
					RichTextLinks []interface{} `json:"richTextLinks"`
					Custom        struct {
					} `json:"custom"`
					Subtitle string `json:"subtitle"`
					Style    struct {
						DefaultStyle struct {
							Margin struct {
								Top string `json:"top"`
							} `json:"margin"`
						} `json:"defaultStyle"`
						ModifiedDate   time.Time `json:"modifiedDate"`
						ExposeTemplate bool      `json:"exposeTemplate"`
						Properties     struct {
							Actions struct {
							} `json:"actions"`
						} `json:"properties"`
						ResourceType string `json:"resourceType"`
					} `json:"style"`
					JSONBody struct {
						Type    string `json:"type"`
						Content []struct {
							Type    string `json:"type"`
							Content []struct {
								Text string `json:"text"`
								Type string `json:"type"`
							} `json:"content"`
						} `json:"content"`
					} `json:"jsonBody"`
					Body    string        `json:"body"`
					Title   string        `json:"title"`
					Actions []interface{} `json:"actions"`
				} `json:"properties,omitempty"`
				Properties5 struct {
					SquarishURL  string `json:"squarishURL"`
					AltText      string `json:"altText"`
					PortraitURL  string `json:"portraitURL"`
					LandscapeURL string `json:"landscapeURL"`
					Title        string `json:"title"`
					Portrait     struct {
						View string `json:"view"`
						ID   string `json:"id"`
						Type string `json:"type"`
						URL  string `json:"url"`
					} `json:"portrait"`
					CopyID        string        `json:"copyId"`
					RichTextLinks []interface{} `json:"richTextLinks"`
					Subtitle      string        `json:"subtitle"`
					ColorTheme    string        `json:"colorTheme"`
					Style         struct {
						DefaultStyle struct {
							Margin struct {
								Top string `json:"top"`
							} `json:"margin"`
						} `json:"defaultStyle"`
						ModifiedDate   time.Time `json:"modifiedDate"`
						ExposeTemplate bool      `json:"exposeTemplate"`
						ResourceType   string    `json:"resourceType"`
					} `json:"style"`
					Actions   []interface{} `json:"actions"`
					Landscape struct {
						View string `json:"view"`
						ID   string `json:"id"`
						Type string `json:"type"`
						URL  string `json:"url"`
					} `json:"landscape"`
				} `json:"properties,omitempty"`
				Properties6 struct {
					CopyID        string        `json:"copyId"`
					Product       []interface{} `json:"product"`
					RichTextLinks []interface{} `json:"richTextLinks"`
					Custom        struct {
					} `json:"custom"`
					Subtitle string `json:"subtitle"`
					Style    struct {
						DefaultStyle struct {
							Margin struct {
								Top string `json:"top"`
							} `json:"margin"`
						} `json:"defaultStyle"`
						ModifiedDate   time.Time `json:"modifiedDate"`
						ExposeTemplate bool      `json:"exposeTemplate"`
						Properties     struct {
							Actions struct {
							} `json:"actions"`
						} `json:"properties"`
						ResourceType string `json:"resourceType"`
					} `json:"style"`
					JSONBody struct {
						Type    string `json:"type"`
						Content []struct {
							Type    string `json:"type"`
							Content []struct {
								Text string `json:"text"`
								Type string `json:"type"`
							} `json:"content"`
						} `json:"content"`
					} `json:"jsonBody"`
					Body    string        `json:"body"`
					Title   string        `json:"title"`
					Actions []interface{} `json:"actions"`
				} `json:"properties,omitempty"`
			} `json:"nodes"`
			PayloadType        string        `json:"payloadType"`
			PublishStartDate   time.Time     `json:"publishStartDate"`
			SupportedLanguages []interface{} `json:"supportedLanguages"`
			PublishEndDate     time.Time     `json:"publishEndDate"`
			SubType            string        `json:"subType"`
			Links              struct {
				Self struct {
					Ref string `json:"ref"`
				} `json:"self"`
			} `json:"links"`
			ID           string `json:"id"`
			RelationalID string `json:"relationalId"`
			Properties   struct {
				ProductCard struct {
					Analytics struct {
						HashKey string `json:"hashKey"`
					} `json:"analytics"`
					SubType    string `json:"subType"`
					ID         string `json:"id"`
					Type       string `json:"type"`
					Version    string `json:"version"`
					Properties struct {
						SquarishURL string `json:"squarishURL"`
						PortraitID  string `json:"portraitId"`
						AltText     string `json:"altText"`
						PortraitURL string `json:"portraitURL"`
						Squarish    struct {
							View        string  `json:"view"`
							AspectRatio float64 `json:"aspectRatio"`
							ID          string  `json:"id"`
							Type        string  `json:"type"`
							URL         string  `json:"url"`
						} `json:"squarish"`
						Portrait struct {
							View        string  `json:"view"`
							AspectRatio float64 `json:"aspectRatio"`
							ID          string  `json:"id"`
							Type        string  `json:"type"`
							URL         string  `json:"url"`
						} `json:"portrait"`
						SquarishID string `json:"squarishId"`
					} `json:"properties"`
				} `json:"productCard"`
				Tagging struct {
					TaxonomyLabel      string   `json:"taxonomyLabel"`
					TaxonomyAttributes []string `json:"taxonomyAttributes"`
				} `json:"tagging"`
				Publish struct {
					CollectionGroups []string `json:"collectionGroups"`
					Collections      []string `json:"collections"`
					Countries        []string `json:"countries"`
					PageID           string   `json:"pageId"`
				} `json:"publish"`
				Custom struct {
				} `json:"custom"`
				Subtitle       string        `json:"subtitle"`
				ConsumerLabels []interface{} `json:"consumerLabels"`
				ThreadType     string        `json:"threadType"`
				Title          string        `json:"title"`
				Seo            struct {
					Slug string `json:"slug"`
				} `json:"seo"`
				Products []struct {
					ProductID          string `json:"productId"`
					TaxonomyAttributes []struct {
						Ids          []string `json:"ids"`
						ResourceType string   `json:"resourceType"`
					} `json:"taxonomyAttributes"`
					StyleColor string `json:"styleColor"`
				} `json:"products"`
			} `json:"properties"`
			ResourceType string `json:"resourceType"`
		} `json:"publishedContent"`
		ProductInfo []struct {
			MerchProduct struct {
				ID               string    `json:"id"`
				SnapshotID       string    `json:"snapshotId"`
				ModificationDate time.Time `json:"modificationDate"`
				Status           string    `json:"status"`
				MerchGroup       string    `json:"merchGroup"`
				StyleCode        string    `json:"styleCode"`
				ColorCode        string    `json:"colorCode"`
				StyleColor       string    `json:"styleColor"`
				Pid              string    `json:"pid"`
				CatalogID        string    `json:"catalogId"`
				ProductGroupID   string    `json:"productGroupId"`
				Brand            string    `json:"brand"`
				Channels         []string  `json:"channels"`
				ConsumerChannels []struct {
					ID           string `json:"id"`
					ResourceType string `json:"resourceType"`
				} `json:"consumerChannels"`
				LegacyCatalogIds       []string      `json:"legacyCatalogIds"`
				Genders                []string      `json:"genders"`
				SizeConverterID        string        `json:"sizeConverterId"`
				SizeGuideID            string        `json:"sizeGuideId"`
				ValueAddedServices     []interface{} `json:"valueAddedServices"`
				SportTags              []string      `json:"sportTags"`
				ClassificationConcepts []interface{} `json:"classificationConcepts"`
				TaxonomyAttributes     []struct {
					ResourceType string   `json:"resourceType"`
					Ids          []string `json:"ids"`
				} `json:"taxonomyAttributes"`
				CommerceCountryInclusions []interface{} `json:"commerceCountryInclusions"`
				CommerceCountryExclusions []interface{} `json:"commerceCountryExclusions"`
				AbTestValues              []interface{} `json:"abTestValues"`
				ProductRollup             struct {
					Type string `json:"type"`
					Key  string `json:"key"`
				} `json:"productRollup"`
				QuantityLimit           int           `json:"quantityLimit"`
				StyleType               string        `json:"styleType"`
				ProductType             string        `json:"productType"`
				MainColor               bool          `json:"mainColor"`
				IsImageAvailable        bool          `json:"isImageAvailable"`
				IsCopyAvailable         bool          `json:"isCopyAvailable"`
				IsAttributionApproved   bool          `json:"isAttributionApproved"`
				IsAppleWatch            bool          `json:"isAppleWatch"`
				ExclusiveAccess         bool          `json:"exclusiveAccess"`
				CommercePublishDate     time.Time     `json:"commercePublishDate"`
				CommerceStartDate       time.Time     `json:"commerceStartDate"`
				CommerceEndDate         time.Time     `json:"commerceEndDate"`
				IsPromoExclusionMessage bool          `json:"isPromoExclusionMessage"`
				LimitRetailExperience   []interface{} `json:"limitRetailExperience"`
				ResourceType            string        `json:"resourceType"`
				Links                   struct {
					Self struct {
						Ref string `json:"ref"`
					} `json:"self"`
				} `json:"links"`
				IsCustomsApproved bool `json:"isCustomsApproved"`
			} `json:"merchProduct"`
			MerchPrice struct {
				ID               string        `json:"id"`
				SnapshotID       string        `json:"snapshotId"`
				ProductID        string        `json:"productId"`
				ParentID         string        `json:"parentId"`
				ParentType       string        `json:"parentType"`
				ModificationDate time.Time     `json:"modificationDate"`
				Country          string        `json:"country"`
				Msrp             float64       `json:"msrp"`
				FullPrice        float64       `json:"fullPrice"`
				CurrentPrice     float64       `json:"currentPrice"`
				Currency         string        `json:"currency"`
				Discounted       bool          `json:"discounted"`
				PromoInclusions  []string      `json:"promoInclusions"`
				PromoExclusions  []interface{} `json:"promoExclusions"`
				ResourceType     string        `json:"resourceType"`
				Links            struct {
					Self struct {
						Ref string `json:"ref"`
					} `json:"self"`
				} `json:"links"`
			} `json:"merchPrice"`
			Availability struct {
				ID           string `json:"id"`
				ProductID    string `json:"productId"`
				ResourceType string `json:"resourceType"`
				Links        struct {
					Self struct {
						Ref string `json:"ref"`
					} `json:"self"`
				} `json:"links"`
				Available bool `json:"available"`
			} `json:"availability"`
			ProductContent struct {
				GlobalPid                      string   `json:"globalPid"`
				LangLocale                     string   `json:"langLocale"`
				ColorDescription               string   `json:"colorDescription"`
				Slug                           string   `json:"slug"`
				FullTitle                      string   `json:"fullTitle"`
				Title                          string   `json:"title"`
				Subtitle                       string   `json:"subtitle"`
				DescriptionHeading             string   `json:"descriptionHeading"`
				Description                    string   `json:"description"`
				TechSpec                       string   `json:"techSpec"`
				ManufacturingCountryOfOrigin   string   `json:"manufacturingCountryOfOrigin"`
				ManufacturingCountriesOfOrigin []string `json:"manufacturingCountriesOfOrigin"`
				Colors                         []struct {
					Type string `json:"type"`
					Name string `json:"name"`
					Hex  string `json:"hex"`
				} `json:"colors"`
				BestFor  []interface{} `json:"bestFor"`
				Athletes []interface{} `json:"athletes"`
				Widths   []interface{} `json:"widths"`
			} `json:"productContent"`
			ImageUrls struct {
				ProductImageURL string `json:"productImageUrl"`
			} `json:"imageUrls"`
			ProductUrls struct {
				SizeChartURL string `json:"sizeChartUrl"`
			} `json:"productUrls"`
			Skus          []NikeProductSku `json:"skus"`
			AvailableSkus []struct {
				ID           string `json:"id"`
				ProductID    string `json:"productId"`
				ResourceType string `json:"resourceType"`
				Links        struct {
					Self struct {
						Ref string `json:"ref"`
					} `json:"self"`
				} `json:"links"`
				Available bool   `json:"available"`
				Level     string `json:"level"`
				SkuID     string `json:"skuId"`
			} `json:"availableSkus"`
			RelatedNBY struct {
				Pbid              string `json:"pbid"`
				PrebuildGroupID   string `json:"prebuildGroupId"`
				PropertiesSeoSlug string `json:"propertiesSeoSlug"`
				MerchPrice        struct {
					Msrp          float64 `json:"msrp"`
					FullPrice     float64 `json:"fullPrice"`
					CurrentPrice  float64 `json:"currentPrice"`
					EmployeePrice int     `json:"employeePrice"`
				} `json:"merchPrice"`
			} `json:"relatedNBY"`
		} `json:"productInfo"`
		Search struct {
			ConceptIds []string `json:"conceptIds"`
		} `json:"search"`
		CollectionTermIds []string `json:"collectionTermIds"`
		ResourceType      string   `json:"resourceType"`
		Links             struct {
			Self struct {
				Ref string `json:"ref"`
			} `json:"self"`
		} `json:"links"`
		Collectionsv2 struct {
			GroupedCollectionTermIds struct {
				NAMING_FAILED  []string `json:""`
				Global         []string `json:"global"`
				NAMING_FAILED0 []string `json:""`
			} `json:"groupedCollectionTermIds"`
			CollectionTermIds []string `json:"collectionTermIds"`
		} `json:"collectionsv2"`
	} `json:"objects"`
}

type NikeProductSku struct {
	ID                    string    `json:"id"`
	SnapshotID            string    `json:"snapshotId"`
	ProductID             string    `json:"productId"`
	ParentID              string    `json:"parentId"`
	ParentType            string    `json:"parentType"`
	ModificationDate      time.Time `json:"modificationDate"`
	MerchGroup            string    `json:"merchGroup"`
	StockKeepingUnitID    string    `json:"stockKeepingUnitId"`
	Gtin                  string    `json:"gtin"`
	NikeSize              string    `json:"nikeSize"`
	SizeConversionID      string    `json:"sizeConversionId"`
	CountrySpecifications []struct {
		Country             string `json:"country"`
		LocalizedSize       string `json:"localizedSize"`
		LocalizedSizePrefix string `json:"localizedSizePrefix"`
		TaxInfo             struct {
			Vat float64 `json:"vat"`
		} `json:"taxInfo"`
	} `json:"countrySpecifications"`
	ResourceType string `json:"resourceType"`
	Links        struct {
		Self struct {
			Ref string `json:"ref"`
		} `json:"self"`
	} `json:"links"`
}
