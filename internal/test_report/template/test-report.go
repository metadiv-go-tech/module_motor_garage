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
                <svg version="1.0" xmlns="http://www.w3.org/2000/svg" style="padding-left: 25px;
                    width="50.000000pt" height="40.000000pt" viewBox="0 0 1512.000000 1202.000000"
                    preserveAspectRatio="xMidYMid meet">
                    <g transform="translate(0.000000,1202.000000) scale(0.100000,-0.100000)"
                    fill="#000000" stroke="none">
                    <path d="M0 6010 l0 -6010 7560 0 7560 0 0 6010 0 6010 -7560 0 -7560 0 0
                    -6010z m6835 1800 c962 -52 1781 -258 2500 -629 315 -163 493 -279 853 -557
                    12 -9 12 -14 2 -31 -13 -20 -21 -17 -259 103 -416 208 -721 336 -1112 465
                    -490 162 -990 262 -1491 298 -185 14 -625 14 -783 1 -513 -43 -976 -161 -1421
                    -362 -74 -33 -134 -64 -134 -68 0 -4 35 -39 78 -77 419 -378 833 -590 1817
                    -934 530 -185 511 -170 220 -165 -254 5 -431 24 -677 71 -470 92 -842 217
                    -2353 792 -302 116 -666 252 -808 303 -142 52 -254 96 -250 100 22 18 237 105
                    383 155 1199 408 2377 591 3435 535z m-4795 -726 c331 -38 656 -117 985 -239
                    77 -29 327 -129 555 -222 1231 -504 1822 -697 2595 -848 l130 -25 -155 5
                    c-444 15 -932 102 -1770 317 -195 50 -467 119 -605 154 -599 154 -1014 213
                    -1400 201 -195 -6 -311 -22 -450 -62 -290 -84 -530 -281 -679 -557 -75 -138
                    -121 -289 -162 -533 l-6 -40 -20 45 c-30 68 -67 258 -80 417 -18 219 1 406 76
                    748 39 175 5 306 -119 470 -34 44 -64 85 -68 91 -13 22 296 76 536 94 158 11
                    465 4 637 -16z m5465 -745 c22 -5 306 -19 630 -31 698 -25 2059 -30 2710 -10
                    679 20 1418 26 1639 13 652 -39 1183 -179 1618 -428 93 -52 167 -97 165 -99
                    -1 -2 -108 15 -237 37 -596 102 -642 104 -2500 100 -2731 -6 -2596 -7 -2820
                    13 -285 24 -484 51 -630 87 -134 32 -502 150 -561 180 -19 9 -61 47 -93 83
                    l-58 66 49 0 c26 0 66 -5 88 -11z m-5215 -711 c0 -20 -39 -689 -44 -748 l-4
                    -55 17 40 c9 22 78 205 152 407 l134 368 371 0 371 0 -331 -647 -332 -648
                    -482 -3 c-454 -2 -482 -1 -482 15 0 10 -25 299 -55 643 -30 344 -55 628 -55
                    633 0 4 167 7 370 7 285 0 370 -3 370 -12z m3870 -3 c0 -9 -11 -67 -25 -131
                    -14 -64 -25 -120 -25 -124 0 -5 33 -10 73 -12 l73 -3 -20 -92 -19 -92 -76 -3
                    -76 -3 -77 -360 c-42 -198 -82 -384 -88 -412 l-11 -53 -265 0 c-145 0 -264 2
                    -264 4 0 2 38 183 85 402 47 219 85 403 85 411 0 9 -18 13 -65 13 -74 0 -71
                    -7 -46 105 l17 80 68 3 c79 3 72 -7 107 165 l23 117 263 0 c228 0 263 -2 263
                    -15z m4520 -12 c0 -16 9 -282 20 -593 11 -311 20 -591 20 -622 l0 -58 -366 0
                    -366 0 7 128 c4 70 8 133 10 140 3 9 -21 12 -99 12 l-104 0 -28 -77 c-15 -43
                    -37 -105 -49 -138 l-22 -60 -367 -3 c-202 -1 -366 2 -364 6 1 5 136 298 298
                    650 l295 642 557 0 558 0 0 -27z m2306 10 c-6 -19 -56 -251 -56 -259 0 -2 34
                    -4 76 -4 67 0 75 -2 70 -17 -2 -10 -12 -53 -21 -95 l-17 -78 -73 0 c-61 0 -75
                    -3 -79 -17 -3 -10 -42 -189 -86 -398 -44 -209 -83 -388 -86 -397 -5 -17 -27
                    -18 -270 -18 l-264 0 5 23 c3 12 43 198 89 412 l83 390 -69 3 c-59 2 -68 5
                    -64 20 2 9 12 52 21 95 l16 77 70 0 c38 0 69 2 69 5 0 7 51 248 55 263 3 9 66
                    12 270 12 249 0 266 -1 261 -17z m-7597 -260 c0 -10 -12 -74 -28 -143 l-27
                    -125 -63 -7 c-115 -13 -180 -46 -219 -110 -11 -20 -44 -151 -82 -328 -34 -162
                    -64 -298 -67 -302 -2 -5 -121 -8 -264 -8 -214 0 -259 2 -259 14 0 7 47 235
                    105 506 58 271 105 494 105 496 0 2 116 4 259 4 l258 0 -5 -27 -5 -26 37 23
                    c55 33 112 49 189 49 58 1 67 -1 66 -16z m-1131 -2 c124 -18 182 -68 182 -155
                    0 -54 -129 -654 -152 -709 -25 -59 -81 -108 -151 -133 -47 -16 -93 -19 -444
                    -22 -403 -3 -475 1 -542 36 -39 20 -71 73 -71 115 0 16 31 177 70 358 53 252
                    76 342 97 381 43 79 108 113 248 131 64 8 703 6 763 -2z m3139 -14 c59 -20 99
                    -61 109 -112 4 -22 -7 -97 -33 -222 l-39 -188 -344 -5 -345 -5 -22 -100 c-35
                    -158 -31 -175 42 -175 62 0 71 15 109 183 l6 27 266 0 267 0 -7 -47 c-13 -95
                    -47 -209 -74 -248 -34 -49 -114 -92 -197 -105 -33 -6 -223 -10 -422 -10 -361
                    0 -435 5 -496 37 -39 20 -67 70 -67 120 0 50 123 639 146 702 34 91 112 144
                    238 161 33 4 230 7 436 6 321 -2 382 -5 427 -19z m6818 0 c70 -23 107 -63 112
                    -124 5 -60 -125 -681 -155 -739 -36 -71 -104 -114 -207 -133 -78 -15 -761 -15
                    -841 0 -103 19 -154 68 -154 148 0 37 124 628 145 690 32 96 109 152 235 170
                    36 5 234 8 440 7 317 -2 383 -5 425 -19z m-6053 -172 l3 -185 79 177 c43 98
                    82 181 85 186 4 4 125 6 269 5 l261 -3 -134 -233 c-74 -129 -135 -242 -135
                    -251 0 -10 20 -116 44 -237 24 -120 47 -236 50 -256 l7 -38 -274 0 -274 0 -7
                    38 c-3 20 -6 82 -7 137 -1 55 -4 134 -8 175 l-6 75 -100 -213 -100 -212 -278
                    0 -277 0 57 93 c32 50 106 170 165 266 59 96 108 184 108 195 0 12 -14 111
                    -30 221 -16 110 -30 210 -30 223 l0 22 265 0 265 0 2 -185z m3348 180 c0 -3
                    -34 -162 -75 -355 -41 -193 -75 -360 -75 -373 0 -38 16 -52 59 -52 72 0 73 3
                    156 390 41 193 77 360 80 373 l5 22 265 0 c146 0 265 -3 265 -7 -1 -5 -49
                    -233 -108 -508 l-107 -500 -257 -3 -258 -2 7 32 c5 26 3 30 -9 25 -131 -50
                    -295 -73 -406 -57 -160 22 -232 87 -232 210 0 39 115 618 157 788 l5 22 264 0
                    c145 0 264 -2 264 -5z"/>
                    <path d="M9974 5127 l-81 -212 70 -3 c38 -2 71 -1 73 1 5 5 26 418 22 423 -2
                    2 -40 -92 -84 -209z"/>
                    <path d="M3783 5160 c-12 -5 -26 -17 -31 -27 -6 -10 -35 -137 -66 -283 -66
                    -307 -66 -313 11 -308 35 2 50 9 64 28 12 16 39 122 74 289 52 245 54 265 40
                    287 -16 25 -51 30 -92 14z"/>
                    <path d="M6855 5170 c-34 -13 -49 -42 -69 -132 l-6 -28 79 0 c89 0 87 -2 97
                    84 6 48 4 55 -17 70 -25 17 -52 20 -84 6z"/>
                    <path d="M13668 5159 c-29 -17 -36 -38 -94 -310 -64 -302 -64 -309 10 -309 69
                        0 75 13 134 296 29 137 55 259 57 271 10 46 -60 80 -107 52z"/>
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
                                        CUSTOMER INSTRUCTIONS AND REPAIRS</th>
                                </tr>
                            </thead>
                            <tbody>
                                <tr>
                                    <td style="width: 33%; border: 0; padding: 0;">
                                        <input type="checkbox" style="vertical-align: middle;">
                                        TUNE
                                    </td>
                                    <td style="width: 33%; border: 0; padding: 0;">
                                        <input type="checkbox" style="vertical-align: middle;"> INJECTOR
                                    </td>
                                    <td style="width: 33%; border: 0; padding: 0;">
                                        <input type="checkbox" style="vertical-align: middle;"> DIESEL SERVICE
                                    </td>
                                </tr>
                                <tr>
                                    <td style="width: 33%; border: 0; padding: 0;">
                                        <input type="checkbox" style="vertical-align: middle;"> STANDARD
                                    </td>
                                    <td style="width: 33%; border: 0; padding: 0;">
                                        <input type="checkbox" style="vertical-align: middle;"> TIMING BELT
                                    </td>
                                    <td style="width: 33%; border: 0; padding: 0;">
                                        <input type="checkbox" style="vertical-align: middle;"> DIAGNOSE
                                    </td>
                                </tr>
                                <tr>
                                    <td style="width: 33%; border: 0; padding: 0;">
                                        <input type="checkbox" style="vertical-align: middle;"> MAJOR SERVICE
                                    </td>
                                    <td style="width: 33%; border: 0; padding: 0;">
                                        <input type="checkbox" style="vertical-align: middle;"> EXHAUST
                                    </td>
                                    <td style="width: 33%; border: 0; padding: 0;">
                                        <input type="checkbox" style="vertical-align: middle;"> AUTO SERVICE
                                    </td>
                                </tr>
                                <tr>
                                    <td style="width: 33%; border: 0; padding: 0;">
                                        <input type="checkbox" style="vertical-align: middle;"> MINOR SERVICE
                                    </td>
                                    <td style="width: 33%; border: 0; padding: 0;">
                                        <input type="checkbox" style="vertical-align: middle;"> SUSPENSION
                                    </td>
                                    <td style="width: 33%; border: 0; padding: 0;">
                                        <input type="checkbox" style="vertical-align: middle;"> VEHICLE CHECK
                                    </td>
                                </tr>
                                <tr>
                                    <td style="width: 33%; border: 0; padding: 0;">
                                        <input type="checkbox" style="vertical-align: middle;"> WIPER DISC SCE
                                    </td>
                                    <td style="width: 33%; border: 0; padding: 0;">
                                        <input type="checkbox" style="vertical-align: middle;"> LOGBOOK SERVICE
                                    </td>
                                    <td style="width: 33%; border: 0; padding: 0;">
                                        <input type="checkbox" style="vertical-align: middle;"> CAB OVER SCHG
                                    </td>
                                </tr>
                                <tr>
                                    <td colspan="3" style="border: 0; padding: 0;">
                                        <input type="checkbox" style="vertical-align: middle;"> AUTHORITY TO PROCEED AS
                                        PER TERMS & CONDITIONS OVERLEAF
                                    </td>
                                </tr>
                                <tr>
                                    <td colspan="3"
                                        style="border: 1px solid #ddd; padding: 4px; height: 30px; vertical-align: top;">
                                        <div style="font-weight: bold; text-decoration: underline;">PRIME ITEM OF
                                            CONCERN</div>
                                    </td>
                                </tr>
                            </tbody>
                        </table>
                        <!-- Road Tests -->
                        <table>
                            <thead>
                                <tr>
                                    <th colspan="3" style="padding: 2px; background-color: #333; color: white;">
                                        1. ROAD TEST</th>
                                </tr>
                            </thead>
                            <tbody>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        1. Fit Seat Cover and Floor Mat
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        2. Check Oil, Water and Tyre Condition
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        3. Obvious Vehicle Damage
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        4. Road Test - Max Speed Reached <span
                                            style="border: 1px solid #ddd; padding-left: 25px; padding-right: 25px; margin-left: 25px; margin-right: 10px;"></span>
                                        kph
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        5. Air Conditioning / Climate / Heater Controls
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        6. Air Conditioning Temperature Check (6-9 degrees)
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        7. Handbrake Operation
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        8. Footbrake Operation
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        9. Clutch / Gearbox Operation
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        10. Automatic Transmission Operation
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        11. Starter Inhibitor Switch
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        12. Requires Injector Service
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                            </tbody>
                        </table>
                        <!-- Engine Tune -->
                        <table>
                            <thead>
                                <tr>
                                    <th style="width: 76%; padding: 2px; background-color: #333; color: white;">
                                        2. ENGINE TUNE</th>
                                    <th style="width: 12%; padding: 2px; text-align: center; background-color: #333; color: white;">Before</th>
                                    <th style="width: 12%; padding: 2px; text-align: center; background-color: #333; color: white;">After</th>
                                </tr>
                            </thead>
                            <tbody>
                                <tr>
                                    <td style="width: 76%; border: 1px solid #ddd; padding: 2px;">
                                        13. Battery Load Test
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 76%; border: 1px solid #ddd; padding: 2px;">
                                        14. Cranking Voltage
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 76%; border: 1px solid #ddd; padding: 2px;">
                                        15. Charging Voltage
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 76%; border: 1px solid #ddd; padding: 2px;">
                                        16. Ignition Timing - Check and Adjust Where Applicable
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 76%; border: 1px solid #ddd; padding: 2px;">
                                        17. Coil / Condenser
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 76%; border: 1px solid #ddd; padding: 2px;">
                                        18. Manufacturers Idle Speed
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 76%; border: 1px solid #ddd; padding: 2px;">
                                        19. Injectors Secure / Visual Leaks
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 76%; border: 1px solid #ddd; padding: 2px;">
                                        20. Fuel Lines / Visual Leaks
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 76%; border: 1px solid #ddd; padding: 2px;">
                                        21. HT Leads
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 76%; border: 1px solid #ddd; padding: 2px;">
                                        22. E.G.R. / E.F.E. Valve
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 76%; border: 1px solid #ddd; padding: 2px;">
                                        23. Distributor Cap
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 76%; border: 1px solid #ddd; padding: 2px;">
                                        24. Rotor
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 76%; border: 1px solid #ddd; padding: 2px;">
                                        25. P.C.V. System
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 76%; border: 1px solid #ddd; padding: 2px;">
                                        26. Battery Cables / Terminals <input type="checkbox"
                                            style="height: 8px; width: 8px; margin-left: 20px;"><span
                                            class="status">Levels</span>
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 76%; border: 1px solid #ddd; padding: 2px;">
                                        27. Plugs Replaced
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                            </tbody>
                        </table>
                        <table>
                            <tbody>
                                <tr>
                                    <td rowspan="2"
                                        style="width: 52%; border: 1px solid #ddd; border-top: none; padding: 2px; height: 20px; vertical-align: top;">
                                        28. Compression or Power Balance
                                    </td>
                                    <td style="width: 6%; border: 1px solid #ddd; border-top: none; padding: 2px;">1
                                    </td>
                                    <td style="width: 6%; border: 1px solid #ddd; border-top: none; padding: 2px;">
                                    </td>
                                    <td style="width: 6%; border: 1px solid #ddd; border-top: none; padding: 2px;">2
                                    </td>
                                    <td style="width: 6%; border: 1px solid #ddd; border-top: none; padding: 2px;">
                                    </td>
                                    <td style="width: 6%; border: 1px solid #ddd; border-top: none; padding: 2px;">3
                                    </td>
                                    <td style="width: 6%; border: 1px solid #ddd; border-top: none; padding: 2px;">
                                    </td>
                                    <td style="width: 6%; border: 1px solid #ddd; border-top: none; padding: 2px;">4
                                    </td>
                                    <td style="width: 6%; border: 1px solid #ddd; border-top: none; padding: 2px;">
                                    </td>
                                </tr>
                                <tr>
                                    <td style="width: 6%; border: 1px solid #ddd; border-top: none; padding: 2px;">5
                                    </td>
                                    <td style="width: 6%; border: 1px solid #ddd; border-top: none; padding: 2px;">
                                    </td>
                                    <td style="width: 6%; border: 1px solid #ddd; border-top: none; padding: 2px;">6
                                    </td>
                                    <td style="width: 6%; border: 1px solid #ddd; border-top: none; padding: 2px;">
                                    </td>
                                    <td style="width: 6%; border: 1px solid #ddd; border-top: none; padding: 2px;">7
                                    </td>
                                    <td style="width: 6%; border: 1px solid #ddd; border-top: none; padding: 2px;">
                                    </td>
                                    <td style="width: 6%; border: 1px solid #ddd; border-top: none; padding: 2px;">8
                                    </td>
                                    <td style="width: 6%; border: 1px solid #ddd; border-top: none; padding: 2px;">
                                    </td>
                                </tr>
                            </tbody>
                        </table>
                        <!-- Light Checks -->
                        <table>
                            <thead>
                                <tr>
                                    <th colspan="3" style="padding: 2px; background-color: #333; color: white;">
                                        3. LIGHT CHECKS</th>
                                </tr>
                            </thead>
                            <tbody>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        29. Headlight Operation
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;"><span
                                            class="status">Hi L</span>
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;"><span
                                            class="status">Hi R</span>
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;"><span
                                            class="status">Lo L</span>
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;"><span
                                            class="status">Lo R</span>
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        30. High Beam Indicator
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        31. Park / Tail Lights
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;"><span
                                            class="status">FL</span>
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;"><span
                                            class="status">FR</span>
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;"><span
                                            class="status">RL</span>
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;"><span
                                            class="status">RR</span>
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        32. Turn Signals / Flashing Rate
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        33. Signal Cancellation
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;"><span
                                            class="status">L</span>
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;"><span
                                            class="status">R</span>
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        34. Brake Lights
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;"><span
                                            class="status">L</span>
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;"><span
                                            class="status">R</span>
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;"><span
                                            class="status">High Level</span>
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        35. License Plate Lights
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;"><span
                                            class="status">L</span>
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;"><span
                                            class="status">R</span>
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        36. Reverse Lights
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;"><span
                                            class="status">L</span>
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;"><span
                                            class="status">R</span>
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        37. All Lenses / Condition
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                            </tbody>
                        </table>
                        <!-- Interior Checks -->
                        <table>
                            <thead>
                                <tr>
                                    <th colspan="3" style="padding: 2px; background-color: #333; color: white;">
                                        4. INTERIOR CHECKS</th>
                                </tr>
                            </thead>
                            <tbody>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        38. Instrument Warning Lights
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        39. Engine Check Light
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        40. Instrument Panel Lights
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        41. Interior Lights / Courtesy Lights
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        42. Windscreen Wiper Blades
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;"><span
                                            class="status">L</span>
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;"><span
                                            class="status">R</span>
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;"><span
                                            class="status">REAR</span>
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        43. Windscreen Condition and Visibility
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        44. Mirrors - Internal / External
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        45. Horn Operation
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        46. Seat Belts
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;"><span
                                            class="status">FL</span>
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;"><span
                                            class="status">FR</span>
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;"><span
                                            class="status">RL</span>
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;"><span
                                            class="status">RC</span>
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;"><span
                                            class="status">RR</span>
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        47. Lubricate Door Locks / Check Straps / Hinges / Bonnet Latch
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;"><span
                                            class="status">FL</span>
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;"><span
                                            class="status">FR</span>
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;"><span
                                            class="status">RL</span>
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;"><span
                                            class="status">RR</span>
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;"><span
                                            class="status">BOOT</span>
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        48. Window Operation
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        49. Boot & Tail Gate Operation
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;"><span
                                            class="status">FL</span>
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;"><span
                                            class="status">FR</span>
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;"><span
                                            class="status">RL</span>
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;"><span
                                            class="status">RC</span>
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        50. Pollen Filter
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
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
                                        5. UNDER BODY</th>
                                </tr>
                            </thead>
                            <tbody>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        51. Drain Oil and Replace Sump Plug/Washer
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        52. Replace Oil Filter
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        53. Engine Oil Leaks
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        54. Gearbox Oil Level - Leaks
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        55. Differential Oil Level - Leaks
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        56. Lubricate Suspension - Where Applicable
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        57. Fuel Line / Leaks / Attachments
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        58. Brake Cables / Hoses / Lines
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        59. Engine Mountings
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;"><span
                                            class="status">GEARBOX MOUNTS</span>
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;"><span
                                            class="status">FRONT</span>
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;"><span
                                            class="status">REAR</span>
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                            </tbody>
                        </table>

                        <table>
                            <thead>
                                <tr>
                                    <th colspan="3" style="padding: 2px; background-color: #333; color: white;">
                                        6. EXHAUST SYSTEM CHECKS</th>
                                </tr>
                            </thead>
                            <tbody>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        60. Engine Pipe / Flange
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        61. Mufflers / Resonators / Pipes
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        62. Catalytic Converter - Visual
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        63. Particulate Filter - Visual
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        64. Support / Hangers
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                            </tbody>
                        </table>

                        <table>
                            <thead>
                                <tr>
                                    <th colspan="3" style="padding: 2px; background-color: #333; color: white;">
                                        7. SUSPENSION / STEERING SYSTEM TEST</th>
                                </tr>
                            </thead>
                            <tbody>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        65. Steering Free Play
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        66. Steering Wear / Leaks / Rack Boots
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        67. Tie Rod Ends
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        68. Suspension Bushes
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        69. Upper and Lower Wishbones
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        70. Sway Bar Rubbers
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;"><span
                                            class="status">FRONT</span>
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;"><span
                                            class="status">REAR</span>
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        71. Ball Joints
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        72. Shock Absorbers
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;"><span
                                            class="status">FL</span>
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;"><span
                                            class="status">FR</span>
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;"><span
                                            class="status">RL</span>
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;"><span
                                            class="status">RR</span>
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        73. Tail Shaft / Joints
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        74. Constant Velocity Joints
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;"><span
                                            class="status">FL</span>
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;"><span
                                            class="status">FR</span>
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;"><span
                                            class="status">RL</span>
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;"><span
                                            class="status">RR</span>
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        75. Rear Spring Bushes
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;"><span
                                            class="status">L</span>
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;"><span
                                            class="status">R</span>
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                            </tbody>
                        </table>
                        <table>
                            <thead>
                                <tr>
                                    <th colspan="3" style="padding: 2px; background-color: #333; color: white;">
                                        8. BREAKING SYSTEM TEST</th>
                                </tr>
                            </thead>
                            <tbody>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        76. Tyre Pressure
                                        <span class="status"
                                            style="margin-left: 15px; margin-right: 2px;">FL</span><input
                                            style="height: 6px; width: 16px;">
                                        <span class="status"
                                            style="margin-left: 15px; margin-right: 2px;">FR</span><input
                                            style="height: 6px; width: 16px;">
                                        <span class="status"
                                            style="margin-left: 15px; margin-right: 2px;">RL</span><input
                                            style="height: 6px; width: 16px;">
                                        <span class="status"
                                            style="margin-left: 15px; margin-right: 2px;">RR</span><input
                                            style="height: 6px; width: 16px;">
                                        <span class="status"
                                            style="margin-left: 15px; margin-right: 2px;">SP</span><input
                                            style="height: 6px; width: 16px;">
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        77. Tyre Condition
                                        <span class="status"
                                            style="margin-left: 20px; margin-right: 2px;">FL</span><input
                                            type="checkbox" style="height: 10px; width: 10px;">
                                        <span class="status"
                                            style="margin-left: 20px; margin-right: 2px;">FR</span><input
                                            type="checkbox" style="height: 10px; width: 10px;">
                                        <span class="status"
                                            style="margin-left: 20px; margin-right: 2px;">RL</span><input
                                            type="checkbox" style="height: 10px; width: 10px;">
                                        <span class="status"
                                            style="margin-left: 20px; margin-right: 2px;">RR</span><input
                                            type="checkbox" style="height: 10px; width: 10px;">
                                        <span class="status"
                                            style="margin-left: 20px; margin-right: 2px;">SP</span><input
                                            type="checkbox" style="height: 10px; width: 10px;">
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        78. Wheel Bearings
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;"><span
                                            class="status">FL</span>
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;"><span
                                            class="status">FR</span>
                                        <span style="margin-left: 10px; margin-right: 10spx;">/</span>
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;"><span
                                            class="status">RL</span>
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;"><span
                                            class="status">RR</span>
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        79. Master Cyl and Booster
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;"><span
                                            class="status">Master</span>
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;"><span
                                            class="status">Booster</span>
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                            </tbody>
                        </table>
                        <table>
                            <tbody>
                                <tr>
                                    <td
                                        style="width: 28%; border: 1px solid #ddd; padding: 2px; border-top: none; border-right: none; border-bottom: none;">
                                        80. Brakes
                                    </td>
                                    <td
                                        style="width: 12%; border: 1px solid #ddd; padding: 2px; border: none; text-align: center;">
                                        Manufacturers Spec Size
                                    </td>
                                    <td
                                        style="width: 12%; border: 1px solid #ddd; padding: 2px; border: none; text-align: center;">
                                        Disc or Drum Size
                                    </td>
                                    <td
                                        style="width: 12%; border: 1px solid #ddd; padding: 2px; border: none; text-align: center;">
                                        Pad or Lining % Worn
                                    </td>
                                    <td
                                        style="width: 12%; border: 1px solid #ddd; padding: 2px; border: none; text-align: center;">
                                        Caliper or Cylinder
                                    </td>
                                    <td
                                        style="width: 12%; border: 1px solid #ddd; padding: 2px; border: none; text-align: center;">
                                        Seats
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px; border-top: none;">
                                    </td>
                                </tr>
                                <tr>
                                    <td
                                        style="width: 28%; border: 1px solid #ddd; padding: 2px; border-top: none; border-right: none; border-bottom: none;">
                                        (a). L.H. Front
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px; text-align: center;">
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px; text-align: center;">
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px; text-align: center;">

                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px; text-align: center;">
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px; text-align: center;">
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px; border-top: none;">
                                    </td>
                                </tr>
                                <tr>
                                    <td
                                        style="width: 28%; border: 1px solid #ddd; padding: 2px; border-top: none; border-right: none; border-bottom: none;">
                                        (b). R.H. Front
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px; text-align: center;">
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px; text-align: center;">
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px; text-align: center;">

                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px; text-align: center;">
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px; text-align: center;">
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px; border-top: none;">
                                    </td>
                                </tr>
                                <tr>
                                    <td
                                        style="width: 28%; border: 1px solid #ddd; padding: 2px; border-top: none; border-right: none; border-bottom: none;">
                                        (c). L.H. Rear
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px; text-align: center;">
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px; text-align: center;">
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px; text-align: center;">

                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px; text-align: center;">
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px; text-align: center;">
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px; border-top: none;">
                                    </td>
                                </tr>
                                <tr>
                                    <td
                                        style="width: 28%; border: 1px solid #ddd; padding: 2px; border-top: none; border-right: none;">
                                        (d). R.H. Rear
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px; text-align: center;">
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px; text-align: center;">
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px; text-align: center;">

                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px; text-align: center;">
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px; text-align: center;">
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px; border-top: none;">
                                    </td>
                                </tr>
                            </tbody>
                        </table>
                        <table>
                            <thead>
                                <tr>
                                    <th colspan="3" style="padding: 2px; background-color: #333; color: white;">
                                        9. UNDER THE BONNET TESTS</th>
                                </tr>
                            </thead>
                            <tbody>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        81. Replace Engine Oil
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        82. Battery Load Test
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;"><span
                                            class="status">Electrolyte</span>
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;"><span
                                            class="status">Clamp</span>
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;"><span
                                            class="status">Terminals</span>
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;"><span
                                            class="status">Cables</span>
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        83. Air Cleaner
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        84. Belts
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        85. Cambelt
                                        <span style="margin-left: 20px;">Manuf Spec:</span> <span
                                            style="border: 1px solid #ddd; padding-left: 25px; padding-right: 25px; margin-left: 5px; margin-right: 10px;"></span>
                                        km
                                        <span style="margin-left: 20px;">Date:</span> <span
                                            style="border: 1px solid #ddd; padding-left: 25px; padding-right: 25px; margin-left: 5px; margin-right: 10px;"></span>
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        86. Coolant Condition
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        87. Conduct External Visual Check of Components
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        88. Start Motor - Check Oil Filter for Leaks
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        89. Power Steering / Hydraulic or Electric Fluid and Condition
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        90. Automatic Transmission Oil Level and Condition Check
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        91. Brake / Clutch Fluid Level and Condition / Top Up
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        92. Windscreen Washers - Operation / Top Up / Additive Added
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        93. Fuel Filter - Visual Only
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        94. Pressurise Cooling System
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; padding: 2px;">
                                        95. Bonnet / Boot Latch Operation
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
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
                                            style="border: 1px solid #ddd; padding-left: 25px; padding-right: 25px; margin-left: 5px; margin-right: 10px;"></span>
                                        kph
                                    </td>
                                    <td style="width: 12%; border: 1px solid #ddd; padding: 2px;"></td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; border-right: none; padding: 2px;">
                                        97. Park Vehicle - Facing Out
                                    </td>
                                    <td
                                        style="width: 12%; border: none; border-right: 1px solid #ddd; border-bottom: 1px solid #ddd; padding: 2px;">
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;">
                                    </td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; border-right: none; padding: 2px;">
                                        98. Reset Service Interval
                                    </td>
                                    <td
                                        style="width: 12%; border: none; border-right: 1px solid #ddd; border-bottom: 1px solid #ddd; padding: 2px;">
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;">
                                    </td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; border-right: none; padding: 2px;">
                                        99. Gloss Tyres
                                    </td>
                                    <td
                                        style="width: 12%; border: none; border-right: 1px solid #ddd; border-bottom: 1px solid #ddd; padding: 2px;">
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;">
                                    </td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; border-right: none; padding: 2px;">
                                        100. Vacuum Carpets Front & Rear and Deodorise Interior
                                    </td>
                                    <td
                                        style="width: 12%; border: none; border-right: 1px solid #ddd; border-bottom: 1px solid #ddd; padding: 2px;">
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;">
                                    </td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; border-right: none; padding: 2px;">
                                        101. Deodorise Interior
                                    </td>
                                    <td
                                        style="width: 12%; border: none; border-right: 1px solid #ddd; border-bottom: 1px solid #ddd; padding: 2px;">
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;">
                                    </td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; border-right: none; padding: 2px;">
                                        102. Wipe Over Dash
                                    </td>
                                    <td
                                        style="width: 12%; border: none; border-right: 1px solid #ddd; border-bottom: 1px solid #ddd; padding: 2px;">
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;">
                                    </td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; border-right: none; padding: 2px;">
                                        103. Clean Windows
                                    </td>
                                    <td
                                        style="width: 12%; border: none; border-right: 1px solid #ddd; border-bottom: 1px solid #ddd; padding: 2px;">
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;">
                                    </td>
                                </tr>
                                <tr>
                                    <td style="width: 88%; border: 1px solid #ddd; border-right: none; padding: 2px;">
                                        104. Remove Seat Cover and Floor Mat
                                    </td>
                                    <td
                                        style="width: 12%; border: none; border-right: 1px solid #ddd; border-bottom: 1px solid #ddd; padding: 2px;">
                                        <input type="checkbox" style="height: 8px; width: 8px; margin-left: 20px;">
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
