package entity

import "github.com/metadiv-go-tech/metagin/v2/base"

type MotorGarageTestReport struct {
	base.Model
	base.ModelWorkspace

	InvoiceId uint                `json:"invoice_id"`
	Invoice   *MotorGarageInvoice `json:"invoice"`

	// CUSTOMER INSTRUCTIONS AND REPAIRS
	Tune               bool   `json:"tune"`
	Standard           bool   `json:"standard"`
	MajorService       bool   `json:"major_service"`
	MinorService       bool   `json:"minor_service"`
	VehicleCheck       bool   `json:"vehicle_check"`
	LogbookMajor       bool   `json:"logbook_major"`
	Injector           bool   `json:"injector"`
	TimingBelt         bool   `json:"timing_belt"`
	Exhaust            bool   `json:"exhaust"`
	Suspension         bool   `json:"suspension"`
	WsteDispSchg       bool   `json:"wste_disp_schg"`
	LogbookStd         bool   `json:"logbook_std"`
	DieselService      bool   `json:"diesel_service"`
	Diagnose           bool   `json:"diagnose"`
	AutoService        bool   `json:"auto_service"`
	Cooling            bool   `json:"cooling"`
	CabOverSchg        bool   `json:"cab_over_schg"`
	AuthorityToProcess bool   `json:"authority_to_process"`
	PrimeItemOfConcern string `json:"prime_item_of_concern"`

	// ROAD TEST
	Item1Checked  bool   `json:"item_1_checked"`
	Item2Checked  bool   `json:"item_2_checked"`
	Item3Checked  bool   `json:"item_3_checked"`
	Item4Checked  bool   `json:"item_4_checked"`
	Item4MaxSpeed string `json:"item_4_max_speed"`
	Item5Checked  bool   `json:"item_5_checked"`
	Item6Checked  bool   `json:"item_6_checked"`
	Item7Checked  bool   `json:"item_7_checked"`
	Item8Checked  bool   `json:"item_8_checked"`
	Item9Checked  bool   `json:"item_9_checked"`
	Item10Checked bool   `json:"item_10_checked"`
	Item11Checked bool   `json:"item_11_checked"`
	Item12Checked bool   `json:"item_12_checked"`

	// ENGINE TUNE
	Item13Before string `json:"item_13_before"`
	Item13After  string `json:"item_13_after"`
	Item14Before string `json:"item_14_before"`
	Item14After  string `json:"item_14_after"`
	Item15Before string `json:"item_15_before"`
	Item15After  string `json:"item_15_after"`
	Item16Before string `json:"item_16_before"`
	Item16After  string `json:"item_16_after"`
	Item17Before string `json:"item_17_before"`
	Item17After  string `json:"item_17_after"`
	Item18Before string `json:"item_18_before"`
	Item18After  string `json:"item_18_after"`
	Item19Before string `json:"item_19_before"`
	Item19After  string `json:"item_19_after"`
	Item20Before string `json:"item_20_before"`
	Item20After  string `json:"item_20_after"`
	Item21Before string `json:"item_21_before"`
	Item21After  string `json:"item_21_after"`
	Item22Before string `json:"item_22_before"`
	Item22After  string `json:"item_22_after"`
	Item23Before string `json:"item_23_before"`
	Item23After  string `json:"item_23_after"`
	Item24Before string `json:"item_24_before"`
	Item24After  string `json:"item_24_after"`
	Item25Before string `json:"item_25_before"`
	Item25After  string `json:"item_25_after"`
	Item26Before string `json:"item_26_before"`
	Item26After  string `json:"item_26_after"`
	Item26Levels bool   `json:"item_26_levels"`
	Item27Before string `json:"item_27_before"`
	Item27After  string `json:"item_27_after"`
	Item28Cell1  string `json:"item_28_cell_1"`
	Item28Cell2  string `json:"item_28_cell_2"`
	Item28Cell3  string `json:"item_28_cell_3"`
	Item28Cell4  string `json:"item_28_cell_4"`
	Item28Cell5  string `json:"item_28_cell_5"`
	Item28Cell6  string `json:"item_28_cell_6"`
	Item28Cell7  string `json:"item_28_cell_7"`
	Item28Cell8  string `json:"item_28_cell_8"`

	// LIGHTS CHECK
	Item29Checked   bool `json:"item_29_checked"`
	Item29HiL       bool `json:"item_29_hi_l"`
	Item29HiR       bool `json:"item_29_hi_r"`
	Item29LoL       bool `json:"item_29_lo_l"`
	Item29LoR       bool `json:"item_29_lo_r"`
	Item30Checked   bool `json:"item_30_checked"`
	Item31Checked   bool `json:"item_31_checked"`
	Item31FL        bool `json:"item_31_fl"`
	Item31FR        bool `json:"item_31_fr"`
	Item31RL        bool `json:"item_31_rl"`
	Item31RR        bool `json:"item_31_rr"`
	Item32Checked   bool `json:"item_32_checked"`
	Item33Checked   bool `json:"item_33_checked"`
	Item33L         bool `json:"item_33_l"`
	Item33R         bool `json:"item_33_r"`
	Item34Checked   bool `json:"item_34_checked"`
	Item34HighLevel bool `json:"item_34_high_level"`
	Item34L         bool `json:"item_34_l"`
	Item34R         bool `json:"item_34_r"`
	Item35Checked   bool `json:"item_35_checked"`
	Item35L         bool `json:"item_35_l"`
	Item35R         bool `json:"item_35_r"`
	Item36Checked   bool `json:"item_36_checked"`
	Item36L         bool `json:"item_36_l"`
	Item36R         bool `json:"item_36_r"`
	Item37Checked   bool `json:"item_37_checked"`

	// INTERIOR CHECKS
	Item38Checked bool `json:"item_38_checked"`
	Item39Checked bool `json:"item_39_checked"`
	Item40Checked bool `json:"item_40_checked"`
	Item41Checked bool `json:"item_41_checked"`
	Item42Checked bool `json:"item_42_checked"`
	Item42L       bool `json:"item_42_l"`
	Item42R       bool `json:"item_42_r"`
	Item42Rear    bool `json:"item_42_rear"`
	Item43Checked bool `json:"item_43_checked"`
	Item44Checked bool `json:"item_44_checked"`
	Item45Checked bool `json:"item_45_checked"`
	Item46Checked bool `json:"item_46_checked"`
	Item46FL      bool `json:"item_46_fl"`
	Item46FR      bool `json:"item_46_fr"`
	Item46RL      bool `json:"item_46_rl"`
	Item46RC      bool `json:"item_46_rc"`
	Item46RR      bool `json:"item_46_rr"`
	Item47Checked bool `json:"item_47_checked"`
	Item47FL      bool `json:"item_47_fl"`
	Item47FR      bool `json:"item_47_fr"`
	Item47RL      bool `json:"item_47_rl"`
	Item47RR      bool `json:"item_47_rr"`
	Item47Boot    bool `json:"item_47_boot"`
	Item48Checked bool `json:"item_48_checked"`
	Item48FL      bool `json:"item_48_fl"`
	Item48FR      bool `json:"item_48_fr"`
	Item48RL      bool `json:"item_48_rl"`
	Item48RR      bool `json:"item_48_rr"`
	Item49Checked bool `json:"item_49_checked"`
	Item50Checked bool `json:"item_50_checked"`

	// UNDER BODY
	Item51Checked bool `json:"item_51_checked"`
	Item52Checked bool `json:"item_52_checked"`
	Item53Checked bool `json:"item_53_checked"`
	Item54Checked bool `json:"item_54_checked"`
	Item55Checked bool `json:"item_55_checked"`
	Item56Checked bool `json:"item_56_checked"`
	Item57Checked bool `json:"item_57_checked"`
	Item58Checked bool `json:"item_58_checked"`
	Item59Checked bool `json:"item_59_checked"`
	Item59Gearbox bool `json:"item_59_gearbox"`
	Item59Front   bool `json:"item_59_front"`
	Item59Rear    bool `json:"item_59_rear"`

	// EXHAUST SYSTEM CHECKS
	Item60Checked bool `json:"item_60_checked"`
	Item61Checked bool `json:"item_61_checked"`
	Item62Checked bool `json:"item_62_checked"`
	Item63Checked bool `json:"item_63_checked"`
	Item64Checked bool `json:"item_64_checked"`

	// SUSPENSION / STEERING SYSTEM TEST
	Item65Checked bool `json:"item_65_checked"`
	Item66Checked bool `json:"item_66_checked"`
	Item67Checked bool `json:"item_67_checked"`
	Item68Checked bool `json:"item_68_checked"`
	Item69Checked bool `json:"item_69_checked"`
	Item70Checked bool `json:"item_70_checked"`
	Item70Front   bool `json:"item_70_front"`
	Item70Rear    bool `json:"item_70_rear"`
	Item71Checked bool `json:"item_71_checked"`
	Item72Checked bool `json:"item_72_checked"`
	Item72FL      bool `json:"item_72_fl"`
	Item72FR      bool `json:"item_72_fr"`
	Item72RL      bool `json:"item_72_rl"`
	Item72RR      bool `json:"item_72_rr"`
	Item73Checked bool `json:"item_73_checked"`
	Item74Checked bool `json:"item_74_checked"`
	Item74FL      bool `json:"item_74_fl"`
	Item74FR      bool `json:"item_74_fr"`
	Item74RL      bool `json:"item_74_rl"`
	Item74RR      bool `json:"item_74_rr"`
	Item75Checked bool `json:"item_75_checked"`
	Item75L       bool `json:"item_75_l"`
	Item75R       bool `json:"item_75_r"`

	// BREAKING SYSTEM TEST
	Item76Checked bool   `json:"item_76_checked"`
	Item76FL      string `json:"item_76_fl"`
	Item76FR      string `json:"item_76_fr"`
	Item76RL      string `json:"item_76_rl"`
	Item76RR      string `json:"item_76_rr"`
	Item76SP      string `json:"item_76_sp"`
	Item77Checked bool   `json:"item_77_checked"`
	Item77FL      bool   `json:"item_77_fl"`
	Item77FR      bool   `json:"item_77_fr"`
	Item77RL      bool   `json:"item_77_rl"`
	Item77RR      bool   `json:"item_77_rr"`
	Item77SP      bool   `json:"item_77_sp"`
	Item78Checked bool   `json:"item_78_checked"`
	Item78FL      bool   `json:"item_78_fl"`
	Item78FR      bool   `json:"item_78_fr"`
	Item78RL      bool   `json:"item_78_rl"`
	Item78RR      bool   `json:"item_78_rr"`
	Item79Checked bool   `json:"item_79_checked"`
	Item79Master  bool   `json:"item_79_master"`
	Item79Booster bool   `json:"item_79_booster"`
}
