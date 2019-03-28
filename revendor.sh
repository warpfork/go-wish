#!/bin/bash
set -euo pipefail
mkdir .tmp

git clone https://github.com/google/go-cmp/ .tmp/go-cmp
cp -r .tmp/go-cmp/cmp/ .
cp .tmp/go-cmp/LICENSE cmp
find cmp/ -type f -name \*.go -print0 | xargs -0 sed -i "s#github.com/google/go-cmp/cmp#github.com/warpfork/go-wish/cmp#"

# Apply patch to make go-cmp always look at unexported fields.
patch cmp/compare.go <<EOF
@@ -375,7 +375,7 @@
				vax = makeAddressable(vx)
				vay = makeAddressable(vy)
			}
-			step.mayForce = s.exporters[t]
+			step.mayForce = s.exporters[t] || true
			step.pvx = vax
			step.pvy = vay
			step.field = t.Field(i)
EOF

# Apply patch to make go-cmp use regular spaces.
patch cmp/report_text.go <<EOF
@@ -12,4 +12,2 @@ import (
 	"time"
-
-	"github.com/warpfork/go-wish/cmp/internal/flags"
 )
@@ -21,3 +19,3 @@
 func (n indentMode) appendIndent(b []byte, d diffMode) []byte {
-	if flags.Deterministic || randBool {
+	if true {
 		// Use regular spaces (U+0020).
EOF


git clone https://github.com/pmezard/go-difflib .tmp/go-difflib
cp -r .tmp/go-difflib/difflib/ .
cp .tmp/go-difflib/LICENSE difflib
