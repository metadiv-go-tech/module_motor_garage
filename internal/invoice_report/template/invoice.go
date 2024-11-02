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
                <svg version="1.0" xmlns="http://www.w3.org/2000/svg"
                    viewBox="0 0 1000.000000 302.000000" preserveAspectRatio="xMidYMid meet">
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
