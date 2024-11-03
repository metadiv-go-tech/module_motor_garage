package template

var InvoiceTemplate = `<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>VortexAuto Invoice {{invoice_number}}</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            margin: 20px;
            background-color: #f9f9f9;
        }

        .header {
            border-bottom: 1px solid #333;
            padding-bottom: 20px;
            margin-bottom: 30px;
            display: flex;
            justify-content: space-between;
            align-items: flex-start;
            gap: 20px;
        }

        .company-info,
        .invoice-info,
        .customer-info {
            flex: 1;
        }

        .company-info h1 {
            margin: 0 0 10px 0;
            font-size: 24px;
            color: #333;
        }

        .company-info p {
            margin: 5px 0;
            font-size: 14px;
        }

        .invoice-info {
            text-align: center;
            padding-top: 5px;
        }

        .invoice-info p {
            margin: 8px 0;
            font-size: 14px;
        }

        .customer-info {
            text-align: right;
            padding-top: 5px;
        }

        .customer-info p {
            margin: 8px 0;
            font-size: 14px;
        }

        .customer-info strong,
        .invoice-info strong {
            display: inline-block;
            width: 120px;
            text-align: right;
            padding-right: 10px;
            color: #555;
        }

        .customer-info p,
        .invoice-info p {
            display: flex;
            justify-content: flex-end;
            align-items: center;
            margin: 8px 0;
            font-size: 14px;
        }

        .customer-info p span,
        .invoice-info p span {
            min-width: 200px;
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
            background-color: white;
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

        .table-item-1st {
            width: 150px;
        }

        .table-item-2nd {
            width: 300px;
        }

        .empty-row td {
            height: 25px;
            background-color: white;
        }

        .summary {
            margin-top: 20px;
            text-align: right;
            border-top: 1px solid #333;
            display: flex;
            justify-content: flex-end;
        }

        .summary-table {
            width: 250px;
        }

        .summary table {
            margin-top: 10px;
            border: none;
            background-color: transparent;=
        }

        .summary table td {
            padding: 8px;
            border: none;
            width: 200px;
        }

        .subtotal {
            font-size: 12px;
            color: #333;
            text-align: right;
        }

        .signatures {
            margin-top: 40px;
            font-size: 14px;
        }

        @media (max-width: 600px) {
            div[style*="display: flex"] {
                flex-direction: column;
            }
        }

        div + div {
            margin-top: 10px;
        }

        .company-logo {
            width: 200px;
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
                    width="62.500000pt" height="50.000000pt" viewBox="0 0 1512.000000 1202.000000"
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
                <strong>ABN:</strong> 123456789
            </p>
            <p style="padding-left: 25px;">Unit 4/22 Musgrave Rd, Coopers Plains QLD 4108</p>
            <p style="padding-left: 25px;">PH: +61 499 299 751</p>
        </div>
        <div class="customer-info">
            <div class="invoice-info">
                <p><strong>Invoice No:</strong><span>{{invoice_number}}</span></p>
                <p><strong>Date:</strong><span>{{invoice_date}}</span></p>
            </div>
            <p><strong>Customer:</strong><span>{{customer_name}}</span></p>
            <p><strong>Vehicle:</strong><span>{{vehicle_name}} ({{vehicle_year}})</span></p>
            <p><strong>Odometer:</strong><span>{{vehicle_odometer}} km</span></p>
            <p><strong>Registration:</strong><span>{{vehicle_registration}}</span></p>
        </div>
    </div>

    <div>
        <div>
            <div class="section-title">Jobs</div>
            <table>
                <thead>
                    <tr>
                        <th class="table-item-1st">Job</th>
                        <th class="table-item-2nd">Description</th>
                        <th>Price</th>
                        <th>Sale (GST)</th>
                    </tr>
                </thead>
                <tbody>
                    {{service_items}}
                </tbody>
            </table>
            <div class="subtotal">
                <p>Services Total (GST Included): <strong>{{services_total}}</strong></p>
            </div>
        </div>

        <div style="margin-top: 25px;">
            <div class="section-title">Parts & Items</div>
            <table>
                <thead>
                    <tr>
                        <th class="table-item-1st">Item</th>
                        <th class="table-item-2nd">Description</th>
                        <th>Qty</th>
                        <th>Price</th>
                        <th>Sale (GST)</th>
                    </tr>
                </thead>
                <tbody>
                    {{product_items}}
                </tbody>
            </table>
            <div class="subtotal">
                <p>Parts & Items Total (GST Included): <strong>{{products_total}}</strong></p>
            </div>
        </div>

        <div style="margin-top: 25px;">
            <div class="section-title">Discounts</div>
            <table>
                <thead>
                    <tr>
                        <th class="table-item-1st">Discount</th>
                        <th class="table-item-2nd">Description</th>
                        <th>Amount</th>
                        <th>Percentage (%)</th>
                    </tr>
                </thead>
                <tbody>
                    {{discount_items}}
                </tbody>
            </table>
            <div class="subtotal">
                <p>Discount Amount: <strong>{{discounts_total}}</strong></p>
                <p>Discount Percentage: <strong>{{discounts_percentage}}%</strong></p>
            </div>
        </div>
    </div>

    <div class="summary">
        <div class="summary-table">
            <table>
                <tbody>
                    <tr>
                        <td>Parts Total:</td>
                        <td style="text-align: right;font-size: 14px;">{{products_total}}</td>
                    </tr>
                    <tr>
                        <td>Services Total:</td>
                        <td style="text-align: right;font-size: 14px;">{{services_total}}</td>
                    </tr>
                    <tr>
                        <td>Discount:</td>
                        <td style="text-align: right;font-size: 14px;">{{discounts_amount}}</td>
                    </tr>
                    <tr>
                        <td>Total (GST Included):</td>
                        <td style="text-align: right;font-size: 14px;"><strong>{{invoice_total}}</strong></td>
                    </tr>
                </tbody>
            </table>
        </div>
    </div>

    <div class="signatures">
        <p>Technician(s): BW</p>
        <p>Manager's Signature: _________________</p>
        <p>Customer's Signature: _________________</p>
    </div>
</body>

</html>
`
