package template

var TestReportTemplate = `<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>VortexAuto Test Report {{invoice_number}}</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            margin: 0px;
            background-color: #f9f9f9;
        }

        .header {
            border-bottom: 1px solid #333;
            padding-bottom: 5px;
            margin-bottom: 0;
            display: flex;
            flex-direction: column;
            align-items: flex-start;
        }

        .company-info {
            display: flex;
            flex-direction: row;
            align-items: center;
            gap: 10px;
        }

        .company-info h1 {
            margin: 0;
            font-size: 18px;
        }

        .company-info p {
            margin: 0;
            font-size: 10px;
        }

        .invoice-info,
        .customer-info {
            text-align: left;
            font-size: 10px;
        }

        .invoice-info p,
        .customer-info p {
            margin: 0px 0;
        }

        .customer-info strong,
        .invoice-info strong {
            display: inline-block;
            width: 60px;
            text-align: right;
            padding-right: 5px;
            color: #555;
        }

        .customer-info p span,
        .invoice-info p span {
            text-align: left;
        }

        .section-title {
            background-color: #f2f2aa;
            padding: 5px;
            font-weight: bold;
            text-align: center;
            margin-bottom: 0;
            margin-top: 10px;
            font-size: 12px;
        }

        table {
            width: 100%;
            border-collapse: collapse;
            margin-top: 0;
            margin-bottom: 0;
            background-color: transparent;
            flex: 1;
        }

        th,
        td {
            padding: 8px;
            border: 1px solid #ddd;
            text-align: left;
            font-size: 12px;
        }

        th {
            background-color: white;
        }

        .test-report-table table,
        .test-report-table td,
        .test-report-table th {
            font-size: 8px !important;
            border: 1px solid #ddd;
        }

        .test-report-table table table {
            border: none;
        }

        .test-report-table table table td,
        .test-report-table table table th {
            border: 1px solid #ddd;
        }

        input[type="checkbox"] {
            width: 10px;
            height: 10px;
            border: 1px solid #ddd;
            border-radius: 0px;
            margin: 1px;
            accent-color: #333;
        }

        @media (max-width: 600px) {
            div[style*="display: flex"] {
                flex-direction: column;
            }
        }

        div+div {
            margin-top: 10px;
        }

        .company-logo {
            width: 180px;
            height: auto;
            margin-bottom: 10px;
        }
    </style>
</head>
<body>
    <div class="header">
        <div class="company-info">
            <p class="company-logo">
                <svg style="width:120px; padding: 5px;" xmlns="http://www.w3.org/2000/svg" version="1.0"
                    preserveAspectRatio="xMidYMid meet" viewBox="199.1 1292.44 2938.79 760.06">
                    <metadata>Created by potrace 1.16, written by Peter Selinger 2001-2019</metadata>
                    <g transform="translate(0.000000,3334.000000) scale(0.100000,-0.100000)" fill="#000000"
                        stroke="none">
                        <path
                            d="M13695 20410 c-2051 -64 -4179 -484 -6390 -1262 -152 -54 -607 -247 -628 -268 -5 -4 4 -12 20 -18 743 -264 1327 -481 2578 -959 1603 -612 2281 -862 3082 -1138 1042 -357 1912 -554 2797 -630 335 -29 868 -37 1016 -15 24 3 -150 67 -990 360 -763 266 -1009 356 -1365 497 -988 392 -1700 785 -2316 1279 -130 104 -369 315 -442 391 l-47 49 167 78 c1116 522 2239 799 3503 866 248 13 856 13 1110 0 1927 -98 3680 -569 5662 -1521 243 -117 792 -397 833 -425 30 -22 38 -17 68 38 23 42 26 53 15 63 -33 30 -532 406 -668 503 -584 418 -1431 871 -2205 1180 -1256 502 -2638 800 -4175 902 -481 31 -1167 44 -1625 30z">
                        </path>
                        <path
                            d="M3300 18853 c-243 -14 -534 -45 -790 -84 -124 -19 -510 -97 -519 -105 -2 -1 57 -80 131 -173 196 -249 260 -367 310 -568 17 -71 22 -117 21 -238 0 -159 -2 -173 -72 -471 -47 -197 -82 -397 -108 -604 -26 -211 -26 -824 0 -1025 23 -173 60 -378 93 -510 25 -99 103 -315 115 -315 3 0 14 57 23 128 102 740 357 1304 786 1732 436 436 974 669 1710 740 386 37 962 18 1510 -51 632 -79 1225 -208 2500 -544 2140 -564 3212 -776 4320 -855 213 -15 629 -23 605 -11 -11 5 -105 24 -210 41 -843 140 -1824 397 -2885 755 -807 273 -1600 579 -3105 1195 -352 144 -723 295 -825 335 -921 361 -1819 566 -2735 625 -143 9 -734 11 -875 3z">
                        </path>
                        <path
                            d="M16198 17194 c65 -86 235 -263 287 -299 50 -33 130 -64 470 -177 658 -220 918 -283 1420 -347 490 -63 994 -101 1620 -122 368 -13 1891 -6 2445 11 1344 40 5450 28 6780 -20 411 -15 602 -39 1370 -170 552 -94 680 -114 687 -107 3 3 -13 15 -34 26 -22 11 -152 86 -289 166 -296 174 -490 274 -680 350 -796 322 -1658 523 -2569 600 -345 29 -575 35 -1355 35 -809 0 -1178 -6 -2660 -40 -1419 -33 -1889 -40 -2630 -40 -967 0 -1443 10 -3427 70 -944 29 -976 31 -1100 56 -70 14 -179 28 -242 31 l-114 6 21 -29z">
                        </path>
                        <path
                            d="M3505 15623 c3 -27 57 -642 120 -1368 63 -726 117 -1346 121 -1378 l6 -57 1051 2 1050 3 728 1420 728 1420 -807 3 c-444 1 -811 0 -817 -2 -5 -2 -154 -395 -330 -874 -176 -479 -321 -867 -323 -863 -3 8 82 1498 94 1659 l6 82 -816 0 -817 0 6 -47z">
                        </path>
                        <path
                            d="M12381 15523 c-17 -82 -46 -218 -64 -303 l-33 -155 -152 -3 c-84 -1 -152 -5 -152 -8 0 -7 -80 -387 -86 -406 -5 -17 7 -18 146 -18 l152 0 -5 -22 c-17 -70 -377 -1767 -377 -1777 0 -8 154 -11 583 -11 l582 0 189 893 c104 490 192 898 194 905 3 9 46 12 168 12 90 0 164 2 164 4 0 4 75 357 85 404 l5 22 -165 0 c-154 0 -166 1 -161 18 5 16 126 581 126 588 0 2 -263 4 -584 4 l-583 0 -32 -147z">
                        </path>
                        <path
                            d="M20961 15628 c-10 -24 -303 -661 -650 -1417 -347 -755 -631 -1377 -631 -1382 0 -5 335 -9 804 -9 l803 0 109 310 109 310 218 0 218 0 -6 -67 c-5 -57 -35 -528 -35 -547 0 -3 363 -6 808 -6 l807 0 -3 78 c-1 42 -20 624 -42 1292 -22 668 -42 1275 -46 1348 l-6 132 -1219 0 -1220 0 -18 -42z m1078 -1028 c-10 -217 -18 -423 -18 -457 l-1 -63 -155 0 c-85 0 -155 3 -155 8 0 7 344 911 346 908 1 0 -7 -179 -17 -396z">
                        </path>
                        <path
                            d="M27286 15653 c-3 -10 -26 -117 -51 -238 -26 -121 -53 -250 -61 -287 l-16 -68 -149 0 c-146 0 -149 0 -154 -22 -13 -55 -85 -393 -85 -400 0 -4 67 -9 148 -10 l149 -3 -188 -880 c-104 -484 -188 -890 -189 -902 l0 -23 580 0 c319 0 580 1 580 3 1 1 88 407 193 902 l193 900 165 3 165 2 45 213 45 212 -163 3 c-90 1 -163 5 -163 7 0 4 114 546 124 588 4 16 -29 17 -579 17 -547 0 -584 -1 -589 -17z">
                        </path>
                        <path
                            d="M11652 15100 c-139 -18 -257 -58 -348 -120 -25 -16 -48 -30 -51 -30 -5 0 -1 30 13 84 l6 26 -568 -2 -568 -3 -232 -1087 c-127 -598 -234 -1100 -237 -1117 l-6 -31 584 0 c322 0 585 2 585 4 0 2 63 300 140 661 110 514 146 669 167 706 14 27 54 76 89 109 99 95 197 131 416 153 l122 12 68 315 c37 173 65 316 63 319 -10 9 -171 9 -243 1z">
                        </path>
                        <path
                            d="M7730 15059 c-296 -34 -459 -128 -546 -314 -43 -92 -43 -92 -200 -836 -153 -723 -160 -772 -114 -877 42 -96 129 -152 295 -190 60 -13 187 -16 855 -19 817 -5 919 -1 1077 37 216 51 349 171 417 372 8 25 79 345 157 710 125 585 142 675 143 763 1 91 -2 105 -27 157 -46 94 -126 146 -282 181 -86 20 -122 21 -895 23 -443 1 -839 -2 -880 -7z m829 -456 c54 -41 54 -43 -74 -649 -70 -331 -127 -577 -139 -604 -31 -64 -88 -93 -175 -88 -81 5 -111 33 -111 104 0 57 223 1112 245 1158 9 18 27 44 40 56 54 51 160 62 214 23z">
                        </path>
                        <path
                            d="M14420 15059 c-344 -37 -511 -154 -589 -414 -11 -39 -82 -358 -158 -710 -128 -603 -136 -647 -137 -755 -1 -109 1 -118 28 -168 47 -85 124 -132 270 -165 87 -20 125 -21 889 -24 676 -3 815 -2 916 11 208 27 334 79 431 176 93 94 128 184 195 505 20 94 38 178 41 188 5 16 -28 17 -590 17 -328 0 -596 -3 -596 -6 0 -16 -70 -324 -80 -349 -26 -69 -86 -105 -174 -105 -44 0 -60 5 -84 25 -50 42 -50 64 9 342 l54 253 747 0 c702 0 747 1 752 18 3 9 40 179 82 377 70 334 75 367 72 455 -3 72 -9 105 -25 136 -23 44 -85 104 -131 128 -49 25 -171 54 -274 66 -115 12 -1534 12 -1648 -1z m823 -431 c32 -12 57 -54 57 -95 0 -15 -10 -76 -22 -135 l-22 -108 -168 0 -168 0 5 23 c3 12 14 62 25 112 12 49 32 109 46 132 46 79 148 108 247 71z">
                        </path>
                        <path
                            d="M29285 15059 c-177 -20 -337 -79 -424 -155 -60 -52 -110 -133 -143 -229 -15 -44 -89 -370 -165 -725 -134 -630 -137 -648 -138 -770 0 -119 1 -127 28 -173 48 -81 126 -127 282 -164 62 -15 165 -17 855 -21 820 -4 939 1 1087 39 107 27 194 72 262 135 66 61 100 113 136 208 35 93 307 1382 313 1482 11 199 -81 304 -313 357 -86 20 -122 21 -895 23 -443 1 -841 -2 -885 -7z m809 -443 c42 -18 66 -51 66 -94 0 -59 -230 -1124 -253 -1172 -29 -59 -82 -90 -155 -90 -89 0 -132 34 -132 105 0 56 229 1124 250 1164 23 46 49 70 97 87 49 17 86 17 127 0z">
                        </path>
                        <path
                            d="M16750 15043 c0 -10 34 -247 75 -527 l74 -509 -365 -589 c-200 -323 -364 -591 -364 -593 0 -3 275 -4 612 -3 l611 3 215 457 c119 252 217 456 218 455 1 -1 10 -208 20 -459 l18 -458 608 0 c572 0 608 1 608 18 0 9 -50 269 -110 577 -61 307 -110 567 -110 577 0 10 135 253 300 539 165 287 300 523 300 525 0 2 -264 3 -587 2 l-586 -3 -176 -393 -176 -394 -6 74 c-4 40 -7 218 -8 396 l-1 322 -585 0 c-551 0 -585 -1 -585 -17z">
                        </path>
                        <path
                            d="M24065 15038 c-67 -289 -347 -1645 -351 -1703 -16 -209 72 -369 249 -450 126 -58 204 -70 447 -70 254 0 356 17 588 94 78 26 142 45 142 42 0 -3 -7 -31 -15 -61 -8 -30 -15 -58 -15 -63 0 -4 255 -6 567 -5 l567 3 232 1085 c127 597 234 1100 238 1118 l8 32 -586 0 c-548 0 -585 -1 -590 -17 -3 -10 -80 -369 -172 -798 -91 -429 -174 -796 -185 -817 -28 -54 -72 -81 -141 -86 -34 -3 -71 1 -88 8 -34 14 -60 55 -60 95 0 22 320 1544 336 1598 5 16 -27 17 -580 17 l-586 0 -5 -22z">
                        </path>
                    </g>
                </svg>
            </p>
            <p style="padding-left: 25px;">
                <strong>ABN:</strong> 15 797 051 352
            </p>
            <p style="padding-left: 25px;">Unit 4/22 Musgrave Rd, Coopers Plains QLD 4108</p>
            <p style="padding-left: 25px;">PH: +61 499 299 751</p>
        </div>
        <div class="info-container"
            style="display: flex; justify-content: space-between; align-items: flex-start; width: 100%;">
            <div class="invoice-info" style="flex: 1; text-align: left;">
                <div style="padding: 0; margin: 0;"><strong>Invoice No:</strong><span>{{invoice_number}}</span>
                </div>
                <div style="padding: 0; margin: 0;"><strong>Date:</strong><span>{{invoice_date}}</span></div>
            </div>
            <div class="customer-info"
                style="margin: 0; flex: 1; text-align: left; display: flex; flex-direction: column; justify-content: flex-start;">
                <div style="padding: 0; margin: 0;"><strong>Customer:</strong><span>{{customer_name}}</span></div>
                <div style="padding: 0; margin: 0; margin-bottom: 5px;"><strong>Vehicle:</strong><span>{{vehicle_name}}
                        ({{vehicle_year}})</span></div>
            </div>
            <div class="customer-info" style="margin: 0; flex: 1; text-align: left;">
                <p><strong>Odometer:</strong><span>{{vehicle_odometer}} km</span></p>
                <p><strong>Registration:</strong><span>{{vehicle_registration}}</span></p>
            </div>
        </div>
    </div>

    <div class="test-report-table" style="margin-top: 0px;">
        <table style="border: 0;">
            <tbody style="border: 0;">
                <tr>
                    <!-- Left Column -->
                    <td style="width: 50%; border: 0; padding: 0;">
                        <!-- Customer Instructions and Repairs -->
                        <table>
                            <thead>
                                <tr>
                                    <th colspan="3" style="padding: 2px; background-color: #333; color: white;">
                                        CUSTOMER INSTRUCTIONS AND REPAIRS
                                    </th>
                                </tr>
                            </thead>
                            <tbody>
                                <tr>
                                    <td style="width: 33%; border: 0; padding: 0;">
                                        <input type="checkbox" style="vertical-align: middle;" {{tune}}> TUNE
                                    </td>
                                    <td style="width: 33%; border: 0; padding: 0;">
                                        <input type="checkbox" style="vertical-align: middle;" {{injector}}> INJECTOR
                                    </td>
                                    <td style="width: 33%; border: 0; padding: 0;">
                                        <input type="checkbox" style="vertical-align: middle;" {{diesel_service}}> DIESEL SERVICE
                                    </td>
                                </tr>
                                <tr>
                                    <td style="width: 33%; border: 0; padding: 0;">
                                        <input type="checkbox" style="vertical-align: middle;" {{standard}}> STANDARD
                                    </td>
                                    <td style="width: 33%; border: 0; padding: 0;">
                                        <input type="checkbox" style="vertical-align: middle;" {{timing_belt}}> TIMING BELT
                                    </td>
                                    <td style="width: 33%; border: 0; padding: 0;">
                                        <input type="checkbox" style="vertical-align: middle;" {{diagnose}}> DIAGNOSE
                                    </td>
                                </tr>
                                <tr>
                                    <td style="width: 33%; border: 0; padding: 0;">
                                        <input type="checkbox" style="vertical-align: middle;" {{major_service}}> MAJOR SERVICE
                                    </td>
                                    <td style="width: 33%; border: 0; padding: 0;">
                                        <input type="checkbox" style="vertical-align: middle;" {{exhaust}}> EXHAUST
                                    </td>
                                    <td style="width: 33%; border: 0; padding: 0;">
                                        <input type="checkbox" style="vertical-align: middle;" {{auto_service}}> AUTO SERVICE
                                    </td>
                                </tr>
                                <tr>
                                    <td style="width: 33%; border: 0; padding: 0;">
                                        <input type="checkbox" style="vertical-align: middle;" {{minor_service}}> MINOR SERVICE
                                    </td>
                                    <td style="width: 33%; border: 0; padding: 0;">
                                        <input type="checkbox" style="vertical-align: middle;" {{suspension}}> SUSPENSION
                                    </td>
                                    <td style="width: 33%; border: 0; padding: 0;">
                                        <input type="checkbox" style="vertical-align: middle;" {{vehicle_check}}> VEHICLE CHECK
                                    </td>
                                </tr>
                                <tr>
                                    <td style="width: 33%; border: 0; padding: 0;">
                                        <input type="checkbox" style="vertical-align: middle;" {{wiper_disc_sce}}> WIPER DISC SCE
                                    </td>
                                    <td style="width: 33%; border: 0; padding: 0;">
                                        <input type="checkbox" style="vertical-align: middle;" {{logbook_service}}> LOGBOOK SERVICE
                                    </td>
                                    <td style="width: 33%; border: 0; padding: 0;">
                                        <input type="checkbox" style="vertical-align: middle;" {{cab_over_schg}}> CAB OVER SCHG
                                    </td>
                                </tr>
                                <tr>
                                    <td colspan="3" style="border: 0; padding: 0;">
                                        <input type="checkbox" style="vertical-align: middle;" {{authority_to_proceed}}>
                                        AUTHORITY TO PROCEED AS PER TERMS & CONDITIONS OVERLEAF
                                    </td>
                                </tr>
                                <tr>
                                    <td colspan="3" style="border: 1px solid #ddd; padding: 4px; height: 30px; vertical-align: top;">
                                        <div style="font-weight: bold; text-decoration: underline;">PRIME ITEM OF CONCERN</div>
                                        {{prime_item_of_concern}}
                                    </td>
                                </tr>
                            </tbody>
                        </table>
                        <!-- Road Tests -->
                        <table>
                            <thead>
                                <tr>
                                    <th colspan="3" style="padding: 2px; background-color: #333; color: white;">
                                        1. ROAD TEST
                                    </th>
                                </tr>
                            </thead>
                            <tbody>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        1. Fit Seat Cover and Floor Mat
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;">
                                        {{fit_seat_cover_and_floor_mat}}</td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        2. Check Oil, Water and Tyre Condition
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;">
                                        {{check_oil_water_and_tyre_condition}}</td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        3. Obvious Vehicle Damage
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;">
                                        {{obvious_vehicle_damage}}</td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        4. Road Test - Max Speed Reached
                                        <span
                                            style="border: 1px solid #ddd; padding-left: 25px; padding-right: 25px; margin-left: 25px; margin-right: 10px;">{{item_10_road_test_max_speed}}</span>
                                        kph
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;">
                                        {{item_10_road_test_max_speed_reached}}</td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        5. Air Conditioning / Climate / Heater Controls
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;">
                                        {{air_conditioning_climate_heater_controls}}</td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        6. Air Conditioning Temperature Check (6-9 degrees)
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;">
                                        {{air_conditioning_temperature_check}}</td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        7. Handbrake Operation
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;">
                                        {{handbrake_operation}}</td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        8. Footbrake Operation
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;">
                                        {{footbrake_operation}}</td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        9. Clutch / Gearbox Operation
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;">
                                        {{clutch_gearbox_operation}}</td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        10. Automatic Transmission Operation
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;">
                                        {{automatic_transmission_operation}}</td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        11. Starter Inhibitor Switch
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;">
                                        {{starter_inhibitor_switch}}</td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        12. Requires Injector Service
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;">
                                        {{requires_injector_service}}</td>
                                </tr>
                            </tbody>
                        </table>
                        <!-- Engine Tune -->
                        <table>
                            <thead>
                                <tr>
                                    <th style="width: 76%; padding: 2px; background-color: #333; color: white">
                                        2. ENGINE TUNE
                                    </th>
                                    <th style="
                                    width: 12%;
                                    padding: 2px;
                                    text-align: center;
                                    background-color: #333;
                                    color: white;
                                    ">
                                        Before
                                    </th>
                                    <th style="
                                    width: 12%;
                                    padding: 2px;
                                    text-align: center;
                                    background-color: #333;
                                    color: white;
                                    ">
                                        After
                                    </th>
                                </tr>
                            </thead>
                            <tbody>
                                <tr>
                                    <td style="width: 76%; border: 1px solid #ddd; padding: 2px">
                                        13. Battery Load Test
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px">
                                        {{battery_load_test_before}}
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px">
                                        {{battery_load_test_after}}
                                    </td>
                                </tr>
                                <tr>
                                    <td style="width: 76%; border: 1px solid #ddd; padding: 2px">
                                        14. Cranking Voltage
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px">
                                        {{cranking_voltage_before}}
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px">
                                        {{cranking_voltage_after}}
                                    </td>
                                </tr>
                                <tr>
                                    <td style="width: 76%; border: 1px solid #ddd; padding: 2px">
                                        15. Charging Voltage
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px">
                                        {{charging_voltage_before}}
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px">
                                        {{charging_voltage_after}}
                                    </td>
                                </tr>
                                <tr>
                                    <td style="width: 76%; border: 1px solid #ddd; padding: 2px">
                                        16. Ignition Timing - Check and Adjust Where Applicable
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px">
                                        {{ignition_timing_before}}
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px">
                                        {{ignition_timing_after}}
                                    </td>
                                </tr>
                                <tr>
                                    <td style="width: 76%; border: 1px solid #ddd; padding: 2px">
                                        17. Coil / Condenser
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px">
                                        {{coil_condenser_before}}
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px">
                                        {{coil_condenser_after}}
                                    </td>
                                </tr>
                                <tr>
                                    <td style="width: 76%; border: 1px solid #ddd; padding: 2px">
                                        18. Manufacturers Idle Speed
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px">
                                        {{idle_speed_before}}
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px">
                                        {{idle_speed_after}}
                                    </td>
                                </tr>
                                <tr>
                                    <td style="width: 76%; border: 1px solid #ddd; padding: 2px">
                                        19. Injectors Secure / Visual Leaks
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px">
                                        {{injectors_secure_before}}
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px">
                                        {{injectors_secure_after}}
                                    </td>
                                </tr>
                                <tr>
                                    <td style="width: 76%; border: 1px solid #ddd; padding: 2px">
                                        20. Fuel Lines / Visual Leaks
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px">
                                        {{fuel_lines_before}}
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px">
                                        {{fuel_lines_after}}
                                    </td>
                                </tr>
                                <tr>
                                    <td style="width: 76%; border: 1px solid #ddd; padding: 2px">
                                        21. HT Leads
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px">
                                        {{ht_leads_before}}
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px">
                                        {{ht_leads_after}}
                                    </td>
                                </tr>
                                <tr>
                                    <td style="width: 76%; border: 1px solid #ddd; padding: 2px">
                                        22. E.G.R. / E.F.E. Valve
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px">
                                        {{egr_valve_before}}
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px">
                                        {{egr_valve_after}}
                                    </td>
                                </tr>
                                <tr>
                                    <td style="width: 76%; border: 1px solid #ddd; padding: 2px">
                                        23. Distributor Cap
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px">
                                        {{distributor_cap_before}}
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px">
                                        {{distributor_cap_after}}
                                    </td>
                                </tr>
                                <tr>
                                    <td style="width: 76%; border: 1px solid #ddd; padding: 2px">
                                        24. Rotor
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px">
                                        {{rotor_before}}
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px">
                                        {{rotor_after}}
                                    </td>
                                </tr>
                                <tr>
                                    <td style="width: 76%; border: 1px solid #ddd; padding: 2px">
                                        25. P.C.V. System
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px">
                                        {{pcv_system_before}}
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px">
                                        {{pcv_system_after}}
                                    </td>
                                </tr>
                                <tr>
                                    <td style="width: 76%; border: 1px solid #ddd; padding: 2px">
                                        26. Battery Cables / Terminals
                                        <input type="checkbox"
                                            {{battery_cables_levels}}
                                            style="height: 8px; width: 8px; margin-left: 20px" /><span
                                            class="status">Levels</span>
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px">
                                        {{battery_cables_before}}
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px">
                                        {{battery_cables_after}}
                                    </td>
                                </tr>
                                <tr>
                                    <td style="width: 76%; border: 1px solid #ddd; padding: 2px">
                                        27. Plugs Replaced
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px">
                                        {{plugs_replaced_before}}
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px">
                                        {{plugs_replaced_after}}
                                    </td>
                                </tr>
                            </tbody>
                        </table>
                        <table>
                            <tbody>
                                <tr>
                                    <td rowspan="2" style="
                                    width: 52%;
                                    border: 1px solid #ddd;
                                    border-top: none;
                                    padding: 2px;
                                    height: 20px;
                                    vertical-align: top;
                                    ">
                                        28. Compression or Power Balance
                                    </td>
                                    <td style="
                                    width: 6%;
                                    border: 1px solid #ddd;
                                    border-top: none;
                                    padding: 2px;
                                    ">
                                        1
                                    </td>
                                    <td style="
                                    width: 6%;
                                    border: 1px solid #ddd;
                                    border-top: none;
                                    padding: 2px;
                                    ">
                                        {{compression_1}}
                                    </td>
                                    <td style="
                                    width: 6%;
                                    border: 1px solid #ddd;
                                    border-top: none;
                                    padding: 2px;
                                    ">
                                        2
                                    </td>
                                    <td style="
                                    width: 6%;
                                    border: 1px solid #ddd;
                                    border-top: none;
                                    padding: 2px;
                                    ">
                                        {{compression_2}}
                                    </td>
                                    <td style="
                                    width: 6%;
                                    border: 1px solid #ddd;
                                    border-top: none;
                                    padding: 2px;
                                    ">
                                        3
                                    </td>
                                    <td style="
                                    width: 6%;
                                    border: 1px solid #ddd;
                                    border-top: none;
                                    padding: 2px;
                                    ">
                                        {{compression_3}}
                                    </td>
                                    <td style="
                                    width: 6%;
                                    border: 1px solid #ddd;
                                    border-top: none;
                                    padding: 2px;
                                    ">
                                        4
                                    </td>
                                    <td style="
                                    width: 6%;
                                    border: 1px solid #ddd;
                                    border-top: none;
                                    padding: 2px;
                                    ">
                                        {{compression_4}}
                                    </td>
                                </tr>
                                <tr>
                                    <td style="
                                    width: 6%;
                                    border: 1px solid #ddd;
                                    border-top: none;
                                    padding: 2px;
                                    ">
                                        5
                                    </td>
                                    <td style="
                                    width: 6%;
                                    border: 1px solid #ddd;
                                    border-top: none;
                                    padding: 2px;
                                    ">
                                        {{compression_5}}
                                    </td>
                                    <td style="
                                    width: 6%;
                                    border: 1px solid #ddd;
                                    border-top: none;
                                    padding: 2px;
                                    ">
                                        6
                                    </td>
                                    <td style="
                                    width: 6%;
                                    border: 1px solid #ddd;
                                    border-top: none;
                                    padding: 2px;
                                    ">
                                        {{compression_6}}
                                    </td>
                                    <td style="
                                    width: 6%;
                                    border: 1px solid #ddd;
                                    border-top: none;
                                    padding: 2px;
                                    ">
                                        7
                                    </td>
                                    <td style="
                                    width: 6%;
                                    border: 1px solid #ddd;
                                    border-top: none;
                                    padding: 2px;
                                    ">
                                        {{compression_7}}
                                    </td>
                                    <td style="
                                    width: 6%;
                                    border: 1px solid #ddd;
                                    border-top: none;
                                    padding: 2px;
                                    ">
                                        8
                                    </td>
                                    <td style="
                                    width: 6%;
                                    border: 1px solid #ddd;
                                    border-top: none;
                                    padding: 2px;
                                    ">
                                        {{compression_8}}
                                    </td>
                                </tr>
                            </tbody>
                        </table>
                        <!-- Light Checks -->
                        <table>
                            <thead>
                                <tr>
                                    <th colspan="3" style="padding: 2px; background-color: #333; color: white">
                                        3. LIGHT CHECKS
                                    </th>
                                </tr>
                            </thead>
                            <tbody>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px">
                                        29. Headlight Operation
                                        <input type="checkbox"
                                            {{headlight_operation_hi_l}}
                                            style="height: 8px; width: 8px; margin-left: 20px" /><span class="status">Hi
                                            L</span>
                                        <input type="checkbox"
                                            {{headlight_operation_hi_r}}
                                            style="height: 8px; width: 8px; margin-left: 20px" /><span class="status">Hi
                                            R</span>
                                        <input type="checkbox"
                                            {{headlight_operation_Lo_l}}
                                            style="height: 8px; width: 8px; margin-left: 20px" /><span class="status">Lo
                                            L</span>
                                        <input type="checkbox"
                                            {{headlight_operation_Lo_r}}
                                            style="height: 8px; width: 8px; margin-left: 20px" /><span class="status">Lo
                                            R</span>
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px">
                                        {{headlight_operation}}
                                    </td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px">
                                        30. High Beam Indicator
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px">
                                        {{high_beam_indicator}}
                                    </td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px">
                                        31. Park / Tail Lights
                                        <input type="checkbox"
                                            {{park_tail_lights_fl}}
                                            style="height: 8px; width: 8px; margin-left: 20px" /><span
                                            class="status">FL</span>
                                        <input type="checkbox"
                                            {{park_tail_lights_fr}}
                                            style="height: 8px; width: 8px; margin-left: 20px" /><span
                                            class="status">FR</span>
                                        <input type="checkbox"
                                            {{park_tail_lights_rl}}
                                            style="height: 8px; width: 8px; margin-left: 20px" /><span
                                            class="status">RL</span>
                                        <input type="checkbox"
                                            {{park_tail_lights_rr}}
                                            style="height: 8px; width: 8px; margin-left: 20px" /><span
                                            class="status">RR</span>
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px">
                                        {{park_tail_lights}}
                                    </td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px">
                                        32. Turn Signals / Flashing Rate
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px">
                                        {{turn_signals_rate}}
                                    </td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px">
                                        33. Signal Cancellation
                                        <input type="checkbox"
                                            {{signal_cancellation_l}}
                                            style="height: 8px; width: 8px; margin-left: 20px" /><span
                                            class="status">L</span>
                                        <input type="checkbox"
                                            {{signal_cancellation_r}}
                                            style="height: 8px; width: 8px; margin-left: 20px" /><span
                                            class="status">R</span>
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px">
                                        {{signal_cancellation}}
                                    </td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px">
                                        34. Brake Lights
                                        <input type="checkbox"
                                            {{brake_lights_l}}
                                            style="height: 8px; width: 8px; margin-left: 20px" /><span
                                            class="status">L</span>
                                        <input type="checkbox"
                                            {{brake_lights_r}}
                                            style="height: 8px; width: 8px; margin-left: 20px" /><span
                                            class="status">R</span>
                                        <input type="checkbox"
                                            {{brake_lights_high_level}}
                                            style="height: 8px; width: 8px; margin-left: 20px" /><span
                                            class="status">High Level</span>
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px">
                                        {{brake_lights}}
                                    </td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px">
                                        35. License Plate Lights
                                        <input type="checkbox"
                                            style="height: 8px; width: 8px; margin-left: 20px" {{license_plate_lights_l}}/><span
                                            class="status">L</span>
                                        <input type="checkbox"
                                            style="height: 8px; width: 8px; margin-left: 20px" {{license_plate_lights_r}}/><span
                                            class="status">R</span>
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px">
                                        {{license_plate_lights}}
                                    </td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px">
                                        36. Reverse Lights
                                        <input type="checkbox"
                                            style="height: 8px; width: 8px; margin-left: 20px" {{reverse_lights_l}}/><span
                                            class="status">L</span>
                                        <input type="checkbox"
                                            style="height: 8px; width: 8px; margin-left: 20px" {{reverse_lights_r}}/><span
                                            class="status">R</span>
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px">
                                        {{reverse_lights}}
                                    </td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px">
                                        37. All Lenses / Condition
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px">
                                        {{all_lenses_condition}}
                                    </td>
                                </tr>
                            </tbody>
                        </table>
                        <!-- Interior Checks -->
                        <table>
                            <thead>
                                <tr>
                                    <th colspan="3" style="padding: 2px; background-color: #333; color: white;">
                                        4. INTERIOR CHECKS
                                    </th>
                                </tr>
                            </thead>
                            <tbody>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        38. Instrument Warning Lights
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;">
                                        {{instrument_warning_lights}}</td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        39. Engine Check Light
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;">{{engine_check_light}}
                                    </td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        40. Instrument Panel Lights
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;">
                                        {{instrument_panel_lights}}</td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        41. Interior Lights / Courtesy Lights
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;">
                                        {{interior_lights_courtesy_lights}}</td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        42. Windscreen Wiper Blades
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;" {{windscreen_wiper_blades_l}}>
                                        <span class="status">L</span>
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;" {{windscreen_wiper_blades_r}}>
                                        <span class="status">R</span>
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;" {{windscreen_wiper_blades_rear}}>
                                        <span class="status">REAR</span>
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;">
                                        {{windscreen_wiper_blades}}</td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        43. Windscreen Condition and Visibility
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;">
                                        {{windscreen_condition_and_visibility}}</td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        44. Mirrors - Internal / External
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;">
                                        {{mirrors_internal_external}}</td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        45. Horn Operation
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;">{{horn_operation}}
                                    </td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        46. Seat Belts
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;" {{seat_belts_fl}}>
                                        <span class="status">FL</span>
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;" {{seat_belts_fr}}>
                                        <span class="status">FR</span>
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;" {{seat_belts_rl}}>
                                        <span class="status">RL</span>
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;" {{seat_belts_rc}}>
                                        <span class="status">RC</span>
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;" {{seat_belts_rr}}>
                                        <span class="status">RR</span>
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;">{{seat_belts}}</td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        47. Lubricate Door Locks / Check Straps / Hinges / Bonnet Latch
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;" {{lubricate_door_locks_check_straps_hinges_bonnet_latch_fl}}>
                                        <span class="status">FL</span>
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;" {{lubricate_door_locks_check_straps_hinges_bonnet_latch_fr}}>
                                        <span class="status">FR</span>
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;" {{lubricate_door_locks_check_straps_hinges_bonnet_latch_rl}}>
                                        <span class="status">RL</span>
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;" {{lubricate_door_locks_check_straps_hinges_bonnet_latch_rr}}>
                                        <span class="status">RR</span>
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;" {{lubricate_door_locks_check_straps_hinges_bonnet_latch_boot}}>
                                        <span class="status">BOOT</span>
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;">
                                        {{lubricate_door_locks_check_straps_hinges_bonnet_latch}}</td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        48. Window Operation
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;" {{window_operation_fl}}>
                                        <span class="status">FL</span>
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;" {{window_operation_fr}}>
                                        <span class="status">FR</span>
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;" {{window_operation_rl}}>
                                        <span class="status">RL</span>
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;" {{window_operation_rr}}>
                                        <span class="status">RR</span>
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;">{{window_operation}}
                                    </td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        49. Boot & Tail Gate Operation
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;">
                                        {{boot_and_tail_gate_operation}}</td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        50. Pollen Filter
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;">{{pollen_filter}}</td>
                                </tr>
                            </tbody>
                        </table>
                    </td>
                        <!-- Right Column -->
                    <td style="width: 50%; border: 0; padding: 0;">
                        <table>
                            <thead>
                                <tr>
                                    <th colspan="3" style="padding: 2px; background-color: #333; color: white;">
                                        5. UNDER BODY
                                    </th>
                                </tr>
                            </thead>
                            <tbody>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        51. Drain Oil and Replace Sump Plug/Washer
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;">
                                        {{drain_oil_replace_sump_plug_washer}}</td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        52. Replace Oil Filter
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;">{{replace_oil_filter}}
                                    </td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        53. Engine Oil Leaks
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;">{{engine_oil_leaks}}
                                    </td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        54. Gearbox Oil Level - Leaks
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;">
                                        {{gearbox_oil_level_leaks}}</td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        55. Differential Oil Level - Leaks
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;">
                                        {{differential_oil_level_leaks}}</td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        56. Lubricate Suspension - Where Applicable
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;">
                                        {{lubricate_suspension_where_applicable}}</td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        57. Fuel Line / Leaks / Attachments
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;">
                                        {{fuel_line_leaks_attachments}}</td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        58. Brake Cables / Hoses / Lines
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;">
                                        {{brake_cables_hoses_lines}}</td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        59. Engine Mountings
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;" {{engine_mountings_gearbox_mounts}}><span
                                            class="status">GEARBOX MOUNTS</span>
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;" {{engine_mountings_front}}><span
                                            class="status">FRONT</span>
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;" {{engine_mountings_rear}}><span
                                            class="status">REAR</span>
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;">{{engine_mountings}}
                                    </td>
                                </tr>
                            </tbody>
                        </table>


                        <table>
                            <thead>
                                <tr>
                                    <th colspan="3" style="padding: 2px; background-color: #333; color: white;">
                                        6. EXHAUST SYSTEM CHECKS
                                    </th>
                                </tr>
                            </thead>
                            <tbody>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        60. Engine Pipe / Flange
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;">{{engine_pipe_flange}}
                                    </td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        61. Mufflers / Resonators / Pipes
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;">
                                        {{mufflers_resonators_pipes}}</td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        62. Catalytic Converter - Visual
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;">
                                        {{catalytic_converter_visual}}</td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        63. Particulate Filter - Visual
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;">
                                        {{particulate_filter_visual}}</td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        64. Support / Hangers
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;">{{support_hangers}}
                                    </td>
                                </tr>
                            </tbody>
                        </table>
                        <table>
                            <thead>
                                <tr>
                                    <th colspan="3" style="padding: 2px; background-color: #333; color: white;">
                                        7. SUSPENSION / STEERING SYSTEM TEST
                                    </th>
                                </tr>
                            </thead>
                            <tbody>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        65. Steering Free Play
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;">{{steering_free_play}}
                                    </td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        66. Steering Wear / Leaks / Rack Boots
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;">
                                        {{steering_wear_leaks_rack_boots}}</td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        67. Tie Rod Ends
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;">{{tie_rod_ends}}</td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        68. Suspension Bushes
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;">{{suspension_bushes}}
                                    </td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        69. Upper and Lower Wishbones
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;">
                                        {{upper_and_lower_wishbones}}</td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        70. Sway Bar Rubbers
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;" {{sway_bar_rubbers_front}}><span
                                            class="status">FRONT</span>
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;" {{sway_bar_rubbers_rear}}><span
                                            class="status">REAR</span>
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;">{{sway_bar_rubbers}}
                                    </td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        71. Ball Joints
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;">{{ball_joints}}</td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        72. Shock Absorbers
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;" {{shock_absorbers_fl}}><span
                                            class="status">FL</span>
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;" {{shock_absorbers_fr}}><span
                                            class="status">FR</span>
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;" {{shock_absorbers_rl}}><span
                                            class="status">RL</span>
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;" {{shock_absorbers_rr}}><span
                                            class="status">RR</span>
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;">{{shock_absorbers}}
                                    </td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        73. Tail Shaft / Joints
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;">{{tail_shaft_joints}}
                                    </td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        74. Constant Velocity Joints
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;" {{constant_velocity_joints_fl}}><span
                                            class="status">FL</span>
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;" {{constant_velocity_joints_fr}}><span
                                            class="status">FR</span>
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;" {{constant_velocity_joints_rl}}><span
                                            class="status">RL</span>
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;" {{constant_velocity_joints_rr}}><span
                                            class="status">RR</span>
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;">
                                        {{constant_velocity_joints}}</td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        75. Rear Spring Bushes
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;" {{rear_spring_bushes_l}}><span
                                            class="status">L</span>
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;" {{rear_spring_bushes_r}}><span
                                            class="status">R</span>
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;">{{rear_spring_bushes}}
                                    </td>
                                </tr>
                            </tbody>
                        </table>
           <table>
                        <thead>
                            <tr>
                                <th colspan="3" style="padding: 2px; background-color: #333; color: white;">
                                    8. BREAKING SYSTEM TEST
                                </th>
                            </tr>
                        </thead>
          <tbody>
                <tr>
                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                        76. Tyre Pressure
                        <span class="status" style="margin-left: 15px; margin-right: 2px;" >FL</span>
                        <input style="height: 6px; width: 16px; font-size: 10px;" value="{{tyre_pressure_status_fl}}">
                        <span class="status" style="margin-left: 15px; margin-right: 2px;" >FR</span>
                        <input style="height: 6px; width: 16px; font-size: 10px;" value="{{tyre_pressure_status_fr}}">
                        <span class="status" style="margin-left: 15px; margin-right: 2px;" >RL</span>
                        <input style="height: 6px; width: 16px; font-size: 10px;" value="{{tyre_pressure_status_rl}}">
                        <span class="status" style="margin-left: 15px; margin-right: 2px;">RR</span>
                        <input style="height: 6px; width: 16px; font-size: 10px;" value="{{tyre_pressure_status_rr}}">
                        <span class="status" style="margin-left: 15px; margin-right: 2px;">SP</span>
                        <input style="height: 6px; width: 16px; font-size: 10px;" value="{{tyre_pressure_status_sp}}">
                    </td>
                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;">{{tyre_pressure_status}}</td>
                </tr>
                <tr>
                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                        77. Tyre Condition
                        <span class="status" style="margin-left: 20px; margin-right: 2px;">FL</span>
                        <input type="checkbox" style="height: 10px; width: 10px;" {{tyre_condition_status_fl}}>
                        <span class="status" style="margin-left: 20px; margin-right: 2px;">FR</span>
                        <input type="checkbox" style="height: 10px; width: 10px;" {{tyre_condition_status_fr}}>
                        <span class="status" style="margin-left: 20px; margin-right: 2px;">RL</span>
                        <input type="checkbox" style="height: 10px; width: 10px;" {{tyre_condition_status_rl}}>
                        <span class="status" style="margin-left: 20px; margin-right: 2px;">RR</span>
                        <input type="checkbox" style="height: 10px; width: 10px;" {{tyre_condition_status_rr}}>
                        <span class="status" style="margin-left: 20px; margin-right: 2px;">SP</span>
                        <input type="checkbox" style="height: 10px; width: 10px;" {{tyre_condition_status_sp}}>
                    </td>
                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;">{{tyre_condition_status}}</td>
                </tr>
                <tr>
                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                        78. Wheel Bearings
                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;"{{wheel_bearings_status_fl}}>
                        <span class="status">FL</span>
                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;"{{wheel_bearings_status_fr}}>
                        <span class="status">FR</span>
                        <span style="margin-left: 10px; margin-right: 10spx;">/</span>
                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;"{{wheel_bearings_status_rl}}>
                        <span class="status">RL</span>
                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;"{{wheel_bearings_status_rr}}>
                        <span class="status">RR</span>
                    </td>
                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;">{{wheel_bearings_status}}</td>
                </tr>
                <tr>
                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                        79. Master Cyl and Booster
                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;" {{master_cyl_booster_status_master}}>
                        <span class="status">Master</span>
                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;" {{master_cyl_booster_status_booster}}>
                        <span class="status">Booster</span>
                    </td>
                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;">{{master_cyl_booster_status}}</td>
                </tr>
            </tbody>
        </table>
        <table>
            <tbody>
                <tr>
                    <td style="width: 28%; border: 1px solid #ddd; padding: 2px;">80. Brakes</td>
                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px; text-align: center;">
                        Manufacturers Spec Size
                    </td>
                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px; text-align: center;">
                        Disc or Drum Size
                    </td>
                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px; text-align: center;">
                        Pad or Lining % Worn
                    </td>
                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px; text-align: center;">
                        Caliper or Cylinder
                    </td>
                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px; text-align: center;">Seats</td>
                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                </tr>
                <tr>
                    <td style="width: 28%; border: 1px solid #ddd; padding: 2px;">(a). L.H. Front</td>
                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px; text-align: center;">
                        {{brakes_lh_front_manufacturers_spec_size}}</td>
                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px; text-align: center;">
                        {{brakes_lh_front_disc_size}}</td>
                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px; text-align: center;">
                        {{brakes_lh_front_pad_worn}}</td>
                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px; text-align: center;">
                        {{brakes_lh_front_caliper}}</td>
                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px; text-align: center;">
                        {{brakes_lh_front_seats}}</td>
                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;">{{brakes_lh_front}}</td>
                </tr>
                <tr>
                    <td
                        style="width: 28%; border: 1px solid #ddd; padding: 2px; border-top: none; border-right: none; border-bottom: none;">
                        (b). R.H. Front
                    </td>
                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px; text-align: center;">
                        {{brakes_rh_front_manufacturers_spec_size}}</td>
                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px; text-align: center;">
                        {{brakes_rh_front_disc_size}}</td>
                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px; text-align: center;">
                        {{brakes_rh_front_pad_worn}}</td>
                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px; text-align: center;">
                        {{brakes_rh_front_caliper}}</td>
                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px; text-align: center;">
                        {{brakes_rh_front_seats}}</td>
                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px; border-top: none;">
                        {{brakes_rh_front}}</td>
                </tr>
                <tr>
                    <td
                        style="width: 28%; border: 1px solid #ddd; padding: 2px; border-top: none; border-right: none; border-bottom: none;">
                        (c). L.H. Rear
                    </td>
                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px; text-align: center;">
                        {{brakes_lh_rear_manufacturers_spec_size}}</td>
                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px; text-align: center;">
                        {{brakes_lh_rear_disc_size}}</td>
                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px; text-align: center;">
                        {{brakes_lh_rear_pad_worn}}</td>
                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px; text-align: center;">
                        {{brakes_lh_rear_caliper}}</td>
                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px; text-align: center;">
                        {{brakes_lh_rear_seats}}</td>
                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px; border-top: none;">
                        {{brakes_lh_rear}}</td>
                </tr>
                <tr>
                    <td style="width: 28%; border: 1px solid #ddd; padding: 2px; border-top: none; border-right: none;">
                        (d). R.H. Rear
                    </td>
                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px; text-align: center;">
                        {{brakes_rh_rear_manufacturers_spec_size}}</td>
                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px; text-align: center;">
                        {{brakes_rh_rear_disc_size}}</td>
                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px; text-align: center;">
                        {{brakes_rh_rear_pad_worn}}</td>
                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px; text-align: center;">
                        {{brakes_rh_rear_caliper}}</td>
                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px; text-align: center;">
                        {{brakes_rh_rear_seats}}</td>
                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px; border-top: none;">
                        {{brakes_rh_rear}}</td>
                </tr>
            </tbody>
        </table>

        <table>
            <thead>
                <tr>
                    <th colspan="3" style="padding: 2px; background-color: #333; color: white;">
                        9. UNDER THE BONNET TESTS
                    </th>
                </tr>
            </thead>
            <tbody>
                <tr>
                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                        81. Replace Engine Oil
                    </td>
                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;">{{replace_engine_oil}}</td>
                </tr>
                <tr>
                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                        82. Battery Load Test
                        <input type="checkbox" {{battery_load_test_electrolyte}} style="height: 8px; width: 8px; margin-left: 20px;"><span
                            class="status">Electrolyte</span>
                        <input type="checkbox" {{battery_load_test_clamp}} style="height: 8px; width: 8px; margin-left: 20px;"><span
                            class="status">Clamp</span>
                        <input type="checkbox" {{battery_load_test_terminals}} style="height: 8px; width: 8px; margin-left: 20px;"><span
                            class="status">Terminals</span>
                        <input type="checkbox" {{battery_load_test_cables}} style="height: 8px; width: 8px; margin-left: 20px;"><span
                            class="status">Cables</span>
                    </td>
                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;">{{battery_load_test}}</td>
                </tr>
                <tr>
                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                        83. Air Cleaner
                    </td>
                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;">{{air_cleaner}}</td>
                </tr>
                <tr>
                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                        84. Belts
                    </td>
                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;">{{belts}}</td>
                </tr>
                <tr>
                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                        85. Cambelt
                        <span style="margin-left: 10px;">Manuf Spec:</span> <span
                            style="border: 1px solid #ddd; padding-left: 20px; padding-right: 20px; margin-left: 5px; margin-right: 10px;">{{cambelt_manuf_spec}}</span>
                        km
                        <span style="margin-left: 10px;">Date:</span> <span
                            style="border: 1px solid #ddd; padding-left: 20px; padding-right: 20px; margin-left: 5px; margin-right: 10px;">{{cambelt_date}}</span>
                    </td>
                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;">{{cambelt}}</td>
                </tr>
                <tr>
                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                        86. Coolant Condition
                    </td>
                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;">{{coolant_condition}}</td>
                </tr>
                <tr>
                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                        87. Conduct External Visual Check of Components
                    </td>
                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;">{{external_visual_check}}</td>
                </tr>
                <tr>
                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                        88. Start Motor - Check Oil Filter for Leaks
                    </td>
                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;">{{start_motor_check_oil_filter}}</td>
                </tr>
                <tr>
                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                        89. Power Steering / Hydraulic or Electric Fluid and Condition
                    </td>
                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;">{{power_steering_condition}}</td>
                </tr>
                <tr>
                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                        90. Automatic Transmission Oil Level and Condition Check
                    </td>
                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;">{{transmission_oil_check}}</td>
                </tr>
                <tr>
                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                        91. Brake / Clutch Fluid Level and Condition / Top Up
                    </td>
                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;">{{brake_clutch_fluid}}</td>
                </tr>
                <tr>
                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                        92. Windscreen Washers - Operation / Top Up / Additive Added
                    </td>
                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;">{{windscreen_washers}}</td>
                </tr>
                <tr>
                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                        93. Fuel Filter - Visual Only
                    </td>
                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;">{{fuel_filter}}</td>
                </tr>
                <tr>
                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                        94. Pressurise Cooling System
                    </td>
                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;">{{pressurise_cooling_system}}</td>
                </tr>
                <tr>
                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                        95. Bonnet / Boot Latch Operation
                    </td>
                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;">{{bonnet_boot_latch}}</td>
                </tr>
            </tbody>
        </table>

        <table>
            <thead>
                <tr>
                    <th colspan="3" style="padding: 2px; background-color: #333; color: white;">
                        10. FINAL PROCEDURES</th>
                </tr>
            </thead>
            <tbody>
                <tr>
                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                        96. Road Test Car (if Roadworthy)
                        <span style="margin-left: 20px;">Max Speed</span> <span
                            style="border: 1px solid #ddd; padding-left: 25px; padding-right: 25px; margin-left: 5px; margin-right: 10px;">{{road_test_max_speed}}</span>
                        kph
                    </td>
                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;">{{road_test_max_speed_reached}}</td>
                </tr>
                <tr>
                    <td style="width: 88%; border: 1px solid #ddd; border-right: none; padding: 2px;">
                        97. Park Vehicle - Facing Out
                    </td>
                    <td
                        style="width: 12%; border: none; border-right: 1px solid #ddd; border-bottom: 1px solid #ddd; padding: 2px;">
                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;" {{park_vehicle_facing_out}}>
                    </td>
                </tr>
                <tr>
                    <td style="width: 88%; border: 1px solid #ddd; border-right: none; padding: 2px;">
                        98. Reset Service Interval
                    </td>
                    <td
                        style="width: 12%; border: none; border-right: 1px solid #ddd; border-bottom: 1px solid #ddd; padding: 2px;">
                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;" {{reset_service_interval}}>
                    </td>
                </tr>
                <tr>
                    <td style="width: 88%; border: 1px solid #ddd; border-right: none; padding: 2px;" >
                        99. Gloss Tyres
                    </td>
                    <td
                        style="width: 12%; border: none; border-right: 1px solid #ddd; border-bottom: 1px solid #ddd; padding: 2px;">
                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;" {{gloss_tyres}}>
                    </td>
                </tr>
                <tr>
                    <td style="width: 88%; border: 1px solid #ddd; border-right: none; padding: 2px;">
                        100. Vacuum Carpets Front & Rear and Deodorise Interior
                    </td>
                    <td
                        style="width: 12%; border: none; border-right: 1px solid #ddd; border-bottom: 1px solid #ddd; padding: 2px;">
                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;" {{vacuum_carpets_and_deodorise}}>
                    </td>
                </tr>
                <tr>
                    <td style="width: 88%; border: 1px solid #ddd; border-right: none; padding: 2px;">
                        101. Deodorise Interior
                    </td>
                    <td
                        style="width: 12%; border: none; border-right: 1px solid #ddd; border-bottom: 1px solid #ddd; padding: 2px;">
                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;" {{deodorise_interior}}>
                    </td>
                </tr>
                <tr>
                    <td style="width: 88%; border: 1px solid #ddd; border-right: none; padding: 2px;">
                        102. Wipe Over Dash
                    </td>
                    <td
                        style="width: 12%; border: none; border-right: 1px solid #ddd; border-bottom: 1px solid #ddd; padding: 2px;">
                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;" {{wipe_over_dash}}>
                    </td>
                </tr>
                <tr>
                    <td style="width: 88%; border: 1px solid #ddd; border-right: none; padding: 2px;">
                        103. Clean Windows
                    </td>
                    <td
                        style="width: 12%; border: none; border-right: 1px solid #ddd; border-bottom: 1px solid #ddd; padding: 2px;">
                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;" {{clean_windows}}>
                    </td>
                </tr>
                <tr>
                    <td style="width: 88%; border: 1px solid #ddd; border-right: none; padding: 2px;">
                        104. Remove Seat Cover and Floor Mat
                    </td>
                    <td
                        style="width: 12%; border: none; border-right: 1px solid #ddd; border-bottom: 1px solid #ddd; padding: 2px;">
                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;" {{remove_seat_cover_and_floor_mat}}>
                   </td>
                                </tr>
                            </tbody>
                        </table>
                    </td>
                </tr>
            </tbody>
        </table>
    </div>
</body>

</html>
`
