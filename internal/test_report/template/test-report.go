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
                <svg version="1.0" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 1000.000000 302.000000"
                    preserveAspectRatio="xMidYMid meet">
                    <metadata>
                        Created by potrace 1.16, written by Peter Selinger 2001-2019
                    </metadata>
                    <g transform="translate(0.000000,302.000000) scale(0.100000,-0.100000)" fill="#000000"
                        stroke="none">
                        <path d="M3880 2643 c-19 -1 -87 -8 -150 -14 -361 -34 -822 -139 -1224 -279
                        -126 -43 -211 -79 -205 -85 3 -2 79 -31 169 -63 91 -33 327 -122 525 -197 820
                        -311 1050 -383 1342 -419 148 -18 313 -27 335 -18 11 5 -64 36 -220 91 -451
                        159 -621 235 -817 366 -103 69 -235 178 -235 194 0 11 157 77 290 122 248 83
                        482 119 776 119 546 0 1055 -127 1622 -407 201 -99 241 -101 115 -6 -439 333
                        -967 529 -1582 587 -148 14 -628 20 -741 9z" />
                        <path d="M1345 2254 c-71 -10 -221 -37 -227 -41 -4 -2 10 -27 31 -56 81 -110
                        88 -149 55 -317 -31 -157 -40 -289 -25 -406 12 -101 38 -197 50 -190 5 3 11
                        29 14 57 14 120 82 270 169 373 54 64 175 141 267 170 184 58 437 57 761 -5
                        57 -11 280 -64 495 -119 436 -110 616 -150 820 -179 143 -20 354 -38 345 -28
                        -3 3 -54 15 -115 27 -332 66 -723 192 -1210 390 -546 222 -646 256 -875 300
                        -90 17 -160 23 -320 25 -113 2 -218 1 -235 -1z" />
                        <path d="M4720 1840 c0 -7 17 -29 38 -49 31 -30 63 -44 187 -84 174 -54 202
                        -60 390 -84 127 -16 233 -18 1175 -15 1276 3 1584 -4 1829 -47 96 -17 176 -29
                        178 -27 4 4 -111 73 -172 103 -110 56 -334 123 -505 152 -237 41 -380 45
                        -1050 32 -344 -6 -664 -13 -711 -16 -173 -9 -1302 25 -1341 40 -12 5 -18 2
                        -18 -5z" />
                        <path d="M1490 1433 c0 -10 5 -58 10 -108 5 -49 19 -200 30 -335 12 -135 24
                        -248 29 -253 4 -4 124 -6 267 -5 l259 3 182 355 182 355 -201 3 c-147 2 -204
                        -1 -209 -9 -4 -7 -40 -100 -79 -208 -40 -108 -75 -198 -78 -201 -5 -5 3 166
                        15 333 l6 87 -207 0 c-190 0 -206 -1 -206 -17z" />
                        <path d="M3735 1375 l-17 -75 -38 0 c-38 0 -38 -1 -49 -50 -6 -28 -11 -53 -11
                        -55 0 -3 16 -5 35 -5 21 0 35 -5 35 -13 0 -7 -20 -106 -45 -220 -25 -115 -43
                        -213 -39 -218 6 -11 264 -12 280 -1 5 4 23 68 38 142 15 74 36 174 48 223 l20
                        87 38 0 c45 0 48 3 56 64 l7 46 -38 0 -38 0 7 38 c3 20 9 54 12 75 l7 37 -146
                        0 -146 0 -16 -75z" />
                        <path d="M5763 1099 c-89 -194 -160 -356 -157 -360 6 -11 375 -12 390 -1 6 4
                        22 39 36 77 l26 70 56 3 57 3 -3 -78 -3 -78 200 0 200 0 -4 175 c-2 96 -7 256
                        -10 355 l-6 180 -310 3 -310 2 -162 -351z" />
                        <path d="M7526 1442 c-3 -5 -10 -38 -17 -75 l-12 -67 -38 0 c-35 0 -39 -3 -48
                        -37 -18 -64 -15 -73 24 -73 19 0 35 -2 35 -5 0 -6 -35 -168 -72 -332 -11 -51
                        -18 -99 -14 -108 4 -12 29 -15 141 -15 134 0 136 0 145 24 5 13 28 116 52 227
                        l43 204 42 3 c41 3 42 4 52 53 6 28 11 52 11 54 0 3 -18 5 -40 5 -37 0 -40 2
                        -35 23 14 63 25 113 25 120 0 10 -288 9 -294 -1z" />
                        <path d="M3500 1294 c-22 -10 -36 -11 -38 -4 -2 6 -56 10 -146 10 l-143 0 -17
                        -82 c-40 -192 -69 -328 -86 -401 -13 -58 -15 -80 -6 -83 7 -3 72 -3 146 -2
                        l134 3 12 45 c7 25 24 101 39 169 14 68 31 131 37 141 17 27 58 47 109 55 l47
                        7 11 61 c7 34 14 70 17 80 5 14 -1 17 -38 17 -24 -1 -59 -8 -78 -16z" />
                        <path d="M2490 1278 c-24 -13 -52 -36 -61 -53 -19 -33 -80 -311 -87 -390 -3
                        -40 0 -54 16 -68 32 -29 125 -39 337 -35 298 5 300 6 358 284 21 100 37 194
                        35 209 -2 16 -15 37 -31 49 -27 20 -43 21 -275 24 -240 3 -248 2 -292 -20z
                        m283 -138 c-2 -74 -51 -273 -70 -287 -23 -18 -41 -16 -51 5 -12 23 49 310 69
                        323 8 5 23 8 34 7 17 -3 20 -10 18 -48z" />
                        <path d="M4202 1287 c-51 -16 -87 -64 -102 -138 -7 -35 -24 -113 -37 -174 -30
                        -136 -32 -181 -7 -206 32 -31 123 -41 339 -37 175 3 205 6 239 23 53 27 73 57
                        90 138 l14 67 -149 0 -148 0 -10 -37 c-6 -21 -11 -44 -11 -51 0 -7 -11 -18
                        -25 -24 -21 -9 -29 -9 -41 4 -14 14 -14 23 -2 79 l14 64 191 3 190 2 23 110
                        c39 186 32 190 -302 189 -132 -1 -247 -6 -266 -12z" />
                        <path d="M4854 1273 c3 -16 12 -76 21 -134 l16 -106 -90 -144 c-49 -79 -87
                        -147 -85 -151 3 -5 70 -8 150 -8 l145 0 57 118 57 117 5 -115 5 -115 144 -3
                        c80 -1 148 1 153 6 5 5 -2 65 -17 138 -14 71 -25 137 -25 147 0 15 88 177 138
                        255 l14 22 -149 0 -149 0 -46 -97 -46 -98 -1 98 -1 97 -151 0 -151 0 6 -27z" />
                        <path d="M6705 1278 c-37 -164 -85 -407 -85 -433 0 -94 130 -140 283 -100 32
                        8 62 15 67 15 6 0 10 -7 10 -15 0 -13 22 -15 134 -15 95 0 136 4 143 13 5 6
                        24 84 43 172 18 88 43 207 56 265 13 58 24 108 24 113 0 4 -66 7 -148 7 l-147
                        0 -37 -177 c-49 -230 -55 -251 -77 -258 -24 -7 -51 4 -51 21 0 8 18 98 40 200
                        22 103 40 192 40 200 0 11 -28 14 -145 14 -142 0 -145 0 -150 -22z" />
                        <path d="M7965 1281 c-23 -10 -49 -30 -58 -45 -24 -35 -102 -405 -93 -440 4
                        -15 20 -33 39 -43 28 -16 62 -18 277 -18 225 0 248 2 285 20 22 11 47 34 56
                        50 20 35 89 350 89 405 0 41 -17 66 -57 81 -13 5 -131 9 -261 9 -208 0 -242
                        -3 -277 -19z m288 -104 c3 -8 -9 -80 -27 -162 -25 -117 -36 -151 -52 -162 -27
                        -17 -54 -7 -54 21 1 41 57 286 70 301 15 19 56 20 63 2z" />
                    </g>
                </svg>
            </p>
            <p style="padding-left: 25px;">
                <strong>ABN:</strong> 123456789
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
