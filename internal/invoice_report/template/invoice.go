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
                <svg style="width:200px; padding: 5px; padding-left: 25px; padding-bottom: 25px;" xmlns="http://www.w3.org/2000/svg" version="1.0" preserveAspectRatio="xMidYMid meet" viewBox="199.1 1292.44 2938.79 760.06"><metadata>Created by potrace 1.16, written by Peter Selinger 2001-2019</metadata><g transform="translate(0.000000,3334.000000) scale(0.100000,-0.100000)" fill="#000000" stroke="none"><path d="M13695 20410 c-2051 -64 -4179 -484 -6390 -1262 -152 -54 -607 -247 -628 -268 -5 -4 4 -12 20 -18 743 -264 1327 -481 2578 -959 1603 -612 2281 -862 3082 -1138 1042 -357 1912 -554 2797 -630 335 -29 868 -37 1016 -15 24 3 -150 67 -990 360 -763 266 -1009 356 -1365 497 -988 392 -1700 785 -2316 1279 -130 104 -369 315 -442 391 l-47 49 167 78 c1116 522 2239 799 3503 866 248 13 856 13 1110 0 1927 -98 3680 -569 5662 -1521 243 -117 792 -397 833 -425 30 -22 38 -17 68 38 23 42 26 53 15 63 -33 30 -532 406 -668 503 -584 418 -1431 871 -2205 1180 -1256 502 -2638 800 -4175 902 -481 31 -1167 44 -1625 30z"></path><path d="M3300 18853 c-243 -14 -534 -45 -790 -84 -124 -19 -510 -97 -519 -105 -2 -1 57 -80 131 -173 196 -249 260 -367 310 -568 17 -71 22 -117 21 -238 0 -159 -2 -173 -72 -471 -47 -197 -82 -397 -108 -604 -26 -211 -26 -824 0 -1025 23 -173 60 -378 93 -510 25 -99 103 -315 115 -315 3 0 14 57 23 128 102 740 357 1304 786 1732 436 436 974 669 1710 740 386 37 962 18 1510 -51 632 -79 1225 -208 2500 -544 2140 -564 3212 -776 4320 -855 213 -15 629 -23 605 -11 -11 5 -105 24 -210 41 -843 140 -1824 397 -2885 755 -807 273 -1600 579 -3105 1195 -352 144 -723 295 -825 335 -921 361 -1819 566 -2735 625 -143 9 -734 11 -875 3z"></path><path d="M16198 17194 c65 -86 235 -263 287 -299 50 -33 130 -64 470 -177 658 -220 918 -283 1420 -347 490 -63 994 -101 1620 -122 368 -13 1891 -6 2445 11 1344 40 5450 28 6780 -20 411 -15 602 -39 1370 -170 552 -94 680 -114 687 -107 3 3 -13 15 -34 26 -22 11 -152 86 -289 166 -296 174 -490 274 -680 350 -796 322 -1658 523 -2569 600 -345 29 -575 35 -1355 35 -809 0 -1178 -6 -2660 -40 -1419 -33 -1889 -40 -2630 -40 -967 0 -1443 10 -3427 70 -944 29 -976 31 -1100 56 -70 14 -179 28 -242 31 l-114 6 21 -29z"></path><path d="M3505 15623 c3 -27 57 -642 120 -1368 63 -726 117 -1346 121 -1378 l6 -57 1051 2 1050 3 728 1420 728 1420 -807 3 c-444 1 -811 0 -817 -2 -5 -2 -154 -395 -330 -874 -176 -479 -321 -867 -323 -863 -3 8 82 1498 94 1659 l6 82 -816 0 -817 0 6 -47z"></path><path d="M12381 15523 c-17 -82 -46 -218 -64 -303 l-33 -155 -152 -3 c-84 -1 -152 -5 -152 -8 0 -7 -80 -387 -86 -406 -5 -17 7 -18 146 -18 l152 0 -5 -22 c-17 -70 -377 -1767 -377 -1777 0 -8 154 -11 583 -11 l582 0 189 893 c104 490 192 898 194 905 3 9 46 12 168 12 90 0 164 2 164 4 0 4 75 357 85 404 l5 22 -165 0 c-154 0 -166 1 -161 18 5 16 126 581 126 588 0 2 -263 4 -584 4 l-583 0 -32 -147z"></path><path d="M20961 15628 c-10 -24 -303 -661 -650 -1417 -347 -755 -631 -1377 -631 -1382 0 -5 335 -9 804 -9 l803 0 109 310 109 310 218 0 218 0 -6 -67 c-5 -57 -35 -528 -35 -547 0 -3 363 -6 808 -6 l807 0 -3 78 c-1 42 -20 624 -42 1292 -22 668 -42 1275 -46 1348 l-6 132 -1219 0 -1220 0 -18 -42z m1078 -1028 c-10 -217 -18 -423 -18 -457 l-1 -63 -155 0 c-85 0 -155 3 -155 8 0 7 344 911 346 908 1 0 -7 -179 -17 -396z"></path><path d="M27286 15653 c-3 -10 -26 -117 -51 -238 -26 -121 -53 -250 -61 -287 l-16 -68 -149 0 c-146 0 -149 0 -154 -22 -13 -55 -85 -393 -85 -400 0 -4 67 -9 148 -10 l149 -3 -188 -880 c-104 -484 -188 -890 -189 -902 l0 -23 580 0 c319 0 580 1 580 3 1 1 88 407 193 902 l193 900 165 3 165 2 45 213 45 212 -163 3 c-90 1 -163 5 -163 7 0 4 114 546 124 588 4 16 -29 17 -579 17 -547 0 -584 -1 -589 -17z"></path><path d="M11652 15100 c-139 -18 -257 -58 -348 -120 -25 -16 -48 -30 -51 -30 -5 0 -1 30 13 84 l6 26 -568 -2 -568 -3 -232 -1087 c-127 -598 -234 -1100 -237 -1117 l-6 -31 584 0 c322 0 585 2 585 4 0 2 63 300 140 661 110 514 146 669 167 706 14 27 54 76 89 109 99 95 197 131 416 153 l122 12 68 315 c37 173 65 316 63 319 -10 9 -171 9 -243 1z"></path><path d="M7730 15059 c-296 -34 -459 -128 -546 -314 -43 -92 -43 -92 -200 -836 -153 -723 -160 -772 -114 -877 42 -96 129 -152 295 -190 60 -13 187 -16 855 -19 817 -5 919 -1 1077 37 216 51 349 171 417 372 8 25 79 345 157 710 125 585 142 675 143 763 1 91 -2 105 -27 157 -46 94 -126 146 -282 181 -86 20 -122 21 -895 23 -443 1 -839 -2 -880 -7z m829 -456 c54 -41 54 -43 -74 -649 -70 -331 -127 -577 -139 -604 -31 -64 -88 -93 -175 -88 -81 5 -111 33 -111 104 0 57 223 1112 245 1158 9 18 27 44 40 56 54 51 160 62 214 23z"></path><path d="M14420 15059 c-344 -37 -511 -154 -589 -414 -11 -39 -82 -358 -158 -710 -128 -603 -136 -647 -137 -755 -1 -109 1 -118 28 -168 47 -85 124 -132 270 -165 87 -20 125 -21 889 -24 676 -3 815 -2 916 11 208 27 334 79 431 176 93 94 128 184 195 505 20 94 38 178 41 188 5 16 -28 17 -590 17 -328 0 -596 -3 -596 -6 0 -16 -70 -324 -80 -349 -26 -69 -86 -105 -174 -105 -44 0 -60 5 -84 25 -50 42 -50 64 9 342 l54 253 747 0 c702 0 747 1 752 18 3 9 40 179 82 377 70 334 75 367 72 455 -3 72 -9 105 -25 136 -23 44 -85 104 -131 128 -49 25 -171 54 -274 66 -115 12 -1534 12 -1648 -1z m823 -431 c32 -12 57 -54 57 -95 0 -15 -10 -76 -22 -135 l-22 -108 -168 0 -168 0 5 23 c3 12 14 62 25 112 12 49 32 109 46 132 46 79 148 108 247 71z"></path><path d="M29285 15059 c-177 -20 -337 -79 -424 -155 -60 -52 -110 -133 -143 -229 -15 -44 -89 -370 -165 -725 -134 -630 -137 -648 -138 -770 0 -119 1 -127 28 -173 48 -81 126 -127 282 -164 62 -15 165 -17 855 -21 820 -4 939 1 1087 39 107 27 194 72 262 135 66 61 100 113 136 208 35 93 307 1382 313 1482 11 199 -81 304 -313 357 -86 20 -122 21 -895 23 -443 1 -841 -2 -885 -7z m809 -443 c42 -18 66 -51 66 -94 0 -59 -230 -1124 -253 -1172 -29 -59 -82 -90 -155 -90 -89 0 -132 34 -132 105 0 56 229 1124 250 1164 23 46 49 70 97 87 49 17 86 17 127 0z"></path><path d="M16750 15043 c0 -10 34 -247 75 -527 l74 -509 -365 -589 c-200 -323 -364 -591 -364 -593 0 -3 275 -4 612 -3 l611 3 215 457 c119 252 217 456 218 455 1 -1 10 -208 20 -459 l18 -458 608 0 c572 0 608 1 608 18 0 9 -50 269 -110 577 -61 307 -110 567 -110 577 0 10 135 253 300 539 165 287 300 523 300 525 0 2 -264 3 -587 2 l-586 -3 -176 -393 -176 -394 -6 74 c-4 40 -7 218 -8 396 l-1 322 -585 0 c-551 0 -585 -1 -585 -17z"></path><path d="M24065 15038 c-67 -289 -347 -1645 -351 -1703 -16 -209 72 -369 249 -450 126 -58 204 -70 447 -70 254 0 356 17 588 94 78 26 142 45 142 42 0 -3 -7 -31 -15 -61 -8 -30 -15 -58 -15 -63 0 -4 255 -6 567 -5 l567 3 232 1085 c127 597 234 1100 238 1118 l8 32 -586 0 c-548 0 -585 -1 -590 -17 -3 -10 -80 -369 -172 -798 -91 -429 -174 -796 -185 -817 -28 -54 -72 -81 -141 -86 -34 -3 -71 1 -88 8 -34 14 -60 55 -60 95 0 22 320 1544 336 1598 5 16 -27 17 -580 17 l-586 0 -5 -22z"></path></g></svg>
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
