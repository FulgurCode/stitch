// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.819
package user

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"fmt"

	"github.com/FulgurCode/stitch/models"
	"github.com/FulgurCode/stitch/view/layout"
)

func Item(product models.Product, stock map[string]int) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Var2 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
			templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
			if !templ_7745c5c3_IsBuffer {
				defer func() {
					templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
					if templ_7745c5c3_Err == nil {
						templ_7745c5c3_Err = templ_7745c5c3_BufErr
					}
				}()
			}
			ctx = templ.InitializeContext(ctx)
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<link rel=\"stylesheet\" href=\"/static/styles/user/item.css\"><script src=\"/static/scripts/shirt-size.js\"></script> ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var3 = []any{Container()}
			templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var3...)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "<div class=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var4 string
			templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var3).String())
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/user/Item.templ`, Line: 1, Col: 0}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = ImageCarousel(product).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "<div class=\"description-container\"><span style=\"gap: 0\"><p class=\"no-line-height\">Shop name</p><h1 style=\"font-weight: 600;\" class=\"no-line-height\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var5 string
			templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(product.Name)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/user/Item.templ`, Line: 20, Col: 82}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, "</h1></span> <span style=\"gap: 2px;\"><p style=\"font-weight: 600;\" class=\"no-line-height\">&#8360;. ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var6 string
			templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("%d", product.Price))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/user/Item.templ`, Line: 23, Col: 109}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, "</p><p>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var7 string
			templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(product.Description)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/user/Item.templ`, Line: 24, Col: 39}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 7, "</p>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templ.RenderScriptItems(ctx, templ_7745c5c3_Buffer, clothingSizeDialog())
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 8, "<p class=\"no-line-height\" onclick=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var8 templ.ComponentScript = clothingSizeDialog()
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var8.Call)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 9, "\">Size:  <button class=\"button-secondary\" style=\"margin-left: var(--padding); padding-inline: var(--padding)\">help</button> <dialog id=\"clothingSizeDialog\"><div class=\"size-recommend\"><form onsubmit=\"event.preventDefault()\" novalidate><section class=\"step active\"><h1 align=\"center\">Basic Details</h1><div class=\"input-group\"><label for=\"gender\">Gender</label> <select id=\"gender\" required><option value=\"\">Select Gender</option> <option value=\"male\">Male</option> <option value=\"female\">Female</option></select></div><div class=\"input-group\"><label for=\"age\">Age</label> <input type=\"number\" id=\"age\" placeholder=\"Age\" required min=\"16\" max=\"80\" step=\"1\"></div><div class=\"input-group\"><label for=\"height\">Height (cm)</label> <input type=\"number\" id=\"height\" placeholder=\"Height (cm)\" required min=\"140\" max=\"200\" step=\"1\"></div><div class=\"input-group\"><label for=\"weight\">Weight (kg)</label> <input type=\"number\" id=\"weight\" placeholder=\"Weight (kg)\" required min=\"30\" max=\"150\" step=\"1\"></div><div class=\"button-container\"><button class=\"button-primary\" onclick=\"nextStep()\">Next</button> ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templ.RenderScriptItems(ctx, templ_7745c5c3_Buffer, closeClothingSizeDialog())
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 10, "<button class=\"button-secondary\" onclick=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var9 templ.ComponentScript = closeClothingSizeDialog()
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var9.Call)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 11, "\">Close</button></div></section><section class=\"step\"><h1 align=\"center\">Physical Attributes</h1><div class=\"input-group\"><label for=\"abdomen\">Abdomen Shape</label> <select id=\"abdomen\" required><option value=\"\">Abdomen Shape</option> <option value=\"flat\">Flat</option> <option value=\"medium\">Medium</option> <option value=\"building\">Building</option></select></div><div class=\"input-group\"><label for=\"chest\">Chest Shape</label> <select id=\"chest\" required><option value=\"\">Chest Shape</option> <option value=\"narrow\">Narrow</option> <option value=\"medium\">Medium</option> <option value=\"wide\">Wide</option></select></div><div class=\"input-group\"><label for=\"fitPreference\">Fit Preference</label> <select id=\"fitPreference\" required><option value=\"\">Fit Preference</option> <option value=\"very_fitted\">Very Fitted</option> <option value=\"fitted\">Fitted</option> <option value=\"normal\">Normal</option> <option value=\"loose\">Loose</option> <option value=\"very_loose\">Very Loose</option></select></div><div class=\"button-container\"><button class=\"button-primary\" onclick=\"handleCalculateSize(); nextStep()\">Next</button> <button class=\"button-secondary\" onclick=\"prevStep()\">Previous</button> ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templ.RenderScriptItems(ctx, templ_7745c5c3_Buffer, closeClothingSizeDialog())
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 12, "<button class=\"button-secondary\" type=\"submit\" onclick=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var10 templ.ComponentScript = closeClothingSizeDialog()
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var10.Call)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 13, "\">Close</button></div></section><section class=\"step\"><div class=\"result\" id=\"result\"><h3 id=\"recommendedSize\">Recommended Size: </h3><p id=\"baseSize\">Base Size: </p></div><div class=\"button-container\"><button class=\"button-primary\" onclick=\"useRecommendedSize(); closeClothingSizeDialog()\">Use Recommended size</button> <button class=\"button-secondary\" onclick=\"prevStep()\">Previous</button> ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templ.RenderScriptItems(ctx, templ_7745c5c3_Buffer, closeClothingSizeDialog())
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 14, "<button class=\"button-secondary\" onclick=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var11 templ.ComponentScript = closeClothingSizeDialog()
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var11.Call)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 15, "\">Close</button></div></section></form></div></dialog></p></span><form name=\"item-order\" hx-on::validation:validate=\"!document.querySelector(&#39;input[name=size]:checked&#39;) &amp;&amp; document.querySelector(&#39;input[name=size]&#39;).reportValidity()\"><div class=\"size-group\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if stock["s"] == 0 {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 16, "<div class=\"radio-button out-of-stock\"><input type=\"radio\" id=\"s\" name=\"size\" value=\"s\" disabled required> <label for=\"small-size\">S</label></div>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 17, "<div class=\"radio-button\"><input type=\"radio\" id=\"s\" name=\"size\" value=\"s\" required> <label for=\"s\">S</label></div>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			if stock["m"] == 0 {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 18, "<div class=\"radio-button out-of-stock\"><input type=\"radio\" id=\"m\" name=\"size\" value=\"m\" disabled required> <label for=\"small-size\">M</label></div>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 19, "<div class=\"radio-button\"><input type=\"radio\" id=\"m\" name=\"size\" value=\"m\" required> <label for=\"m\">M</label></div>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			if stock["l"] == 0 {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 20, "<div class=\"radio-button out-of-stock\"><input type=\"radio\" id=\"l\" name=\"size\" value=\"l\" disabled required> <label for=\"small-size\">L</label></div>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 21, "<div class=\"radio-button\"><input type=\"radio\" id=\"l\" name=\"size\" value=\"l\" required> <label for=\"l\">L</label></div>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			if stock["xl"] == 0 {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 22, "<div class=\"radio-button out-of-stock\"><input type=\"radio\" id=\"xl\" name=\"size\" value=\"xl\" disabled required> <label for=\"small-size\">XL</label></div>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 23, "<div class=\"radio-button\"><input type=\"radio\" id=\"xl\" name=\"size\" value=\"xl\" required> <label for=\"xl\">XL</label></div>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			if stock["xxl"] == 0 {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 24, "<div class=\"radio-button out-of-stock\"><input type=\"radio\" id=\"xxl\" name=\"size\" value=\"xxl\" disabled required> <label for=\"small-size\">XXL</label></div>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 25, "<div class=\"radio-button\"><input type=\"radio\" id=\"xxl\" name=\"size\" value=\"xxl\" required> <label for=\"xxl\">XXL</label></div>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			if stock["xxxl"] == 0 {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 26, "<div class=\"radio-button out-of-stock\"><input type=\"radio\" id=\"xxxl\" name=\"size\" value=\"xxxl\" disabled required> <label for=\"small-size\">XXXL</label></div>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 27, "<div class=\"radio-button\"><input type=\"radio\" id=\"xxxl\" name=\"size\" value=\"xxxl\" required> <label for=\"xxxl\">XXXL</label></div>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 28, "</div><span><button id=\"buy-button\" class=\"button-primary\" hx-include=\"closest form\" hx-get=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var12 string
			templ_7745c5c3_Var12, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("/order/%s", product.Id))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/user/Item.templ`, Line: 186, Col: 68}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var12))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 29, "\" hx-target=\"body\" hx-push-url=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var13 string
			templ_7745c5c3_Var13, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("/order/%s", product.Id))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/user/Item.templ`, Line: 187, Col: 90}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var13))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 30, "\" hx-validate=\"true\">BUY</button> <button id=\"cart-button\" hx-include=\"closest form\" hx-get=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var14 string
			templ_7745c5c3_Var14, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("/add-to-cart/%s", product.Id))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/user/Item.templ`, Line: 193, Col: 74}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var14))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 31, "\" hx-swap=\"none\" hx-on::after-request=\"showToast(&#39;Product Added to your Cart&#39;, &#39;success&#39;)\" class=\"button-secondary\" hx-validate=\"true\">Add to Cart</button></span></form><br></div></div><script defer>\n            function checkAllSizesDisabled() {\n                const sizeInputs = document.querySelectorAll('input[name=\"size\"]');\n                const allDisabled = Array.from(sizeInputs).every(input => input.disabled);\n                \n                const buyButton = document.getElementById('buy-button');\n                const cartButton = document.getElementById('cart-button');\n                \n                if (allDisabled) {\n                    buyButton.disabled = true;\n                    cartButton.disabled = true;\n                    \n                    // Remove HTMX attributes to prevent interactions\n                    buyButton.removeAttribute('hx-get');\n                    buyButton.removeAttribute('hx-target');\n                    buyButton.removeAttribute('hx-push-url');\n                    cartButton.removeAttribute('hx-get');\n                    cartButton.removeAttribute('hx-swap');\n                    \n                    buyButton.classList.add('disabled');\n                    cartButton.classList.add('disabled');\n                }\n            }\n            document.addEventListener('DOMContentLoaded', checkAllSizesDisabled);\n            document.addEventListener('htmx:afterSettle', checkAllSizesDisabled);\n\n            function handleCalculateSize() {\n                const formData = {\n                    gender: document.getElementById('gender').value,\n                    age: parseInt(document.getElementById('age').value),\n                    height: parseInt(document.getElementById('height').value),\n                    weight: parseInt(document.getElementById('weight').value),\n                    abdomen: document.getElementById('abdomen').value,\n                    chest: document.getElementById('chest').value,\n                    fitPreference: document.getElementById('fitPreference').value,\n                };\n\n                const result = calculateSize(formData);\n\n                document.getElementById('result').style.display = 'block';\n                document.getElementById('recommendedSize').innerText = `Recommended Size: ${result.recommendedSize}`;\n                document.getElementById('baseSize').innerText = `Base Size: ${result.baseSize}`;\n            }\n\n            function validateStep(stepIndex) {\n                const steps = document.querySelectorAll('.step');\n                const inputs = steps[stepIndex].querySelectorAll('input, select');\n                let isValid = true;\n\n                for (const input of inputs) {\n                    if (!input.checkValidity()) {\n                        isValid = false;\n                        input.classList.add('invalid');\n                        input.addEventListener('input', () => input.classList.remove('invalid'), { once: true });\n                    }\n                }\n\n                return isValid;\n            }\n\n            function useRecommendedSize() {\n                const recommendedSize = document.getElementById('recommendedSize').textContent.replace('Recommended Size: ', '').trim().toLowerCase();\n\n                const stock = {\n                    s: document.querySelector('#s').disabled ? 0 : 1,\n                    m: document.querySelector('#m').disabled ? 0 : 1,\n                    l: document.querySelector('#l').disabled ? 0 : 1,\n                    xl: document.querySelector('#xl').disabled ? 0 : 1,\n                    xxl: document.querySelector('#xxl').disabled ? 0 : 1,\n                    xxxl: document.querySelector('#xxxl').disabled ? 0 : 1\n                };\n\n                if (stock[recommendedSize] !== 0) {\n                    const sizeRadioButton = document.querySelector(`input[name=\"size\"][value=\"${recommendedSize}\"]`);\n                    if (sizeRadioButton) {\n                        sizeRadioButton.checked = true;\n                    }\n                } else {\n                    showToast(\"Recomended size is not available\", \"error\")\n                }\n            }\n\n            let currentStep = 0;\n\n            function stepReset(){\n                currentStep=0;\n                showStep(0)\n            }\n\n            function showStep(stepIndex) {\n                const steps = document.querySelectorAll('.step');\n                steps.forEach((step, index) => {\n                    step.classList.remove('active');\n                    if (index === stepIndex) {\n                        step.classList.add('active');\n                    }\n                });\n            }\n\n            function nextStep() {\n                const steps = document.querySelectorAll('.step');\n                if (validateStep(currentStep) && currentStep < steps.length - 1) {\n                    currentStep++;\n                    showStep(currentStep);\n                }\n            }\n\n            function prevStep() {\n                if (currentStep > 0) {\n                    currentStep--;\n                    showStep(currentStep);\n                }\n            }\n\n            function closeClothingSizeDialog() {\n                let dialog = document.getElementById(\"clothingSizeDialog\");\n                stepReset();\n                dialog.close();\n            }\n\n                \n        </script>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return nil
		})
		templ_7745c5c3_Err = layout.User().Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

func clothingSizeDialog() templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_clothingSizeDialog_7f75`,
		Function: `function __templ_clothingSizeDialog_7f75(){let dialog = document.getElementById("clothingSizeDialog");
    dialog.showModal();
}`,
		Call:       templ.SafeScript(`__templ_clothingSizeDialog_7f75`),
		CallInline: templ.SafeScriptInline(`__templ_clothingSizeDialog_7f75`),
	}
}

func closeClothingSizeDialog() templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_closeClothingSizeDialog_40a3`,
		Function: `function __templ_closeClothingSizeDialog_40a3(){let dialog = document.getElementById("clothingSizeDialog");
    stepReset();
    dialog.close();
}`,
		Call:       templ.SafeScript(`__templ_closeClothingSizeDialog_40a3`),
		CallInline: templ.SafeScriptInline(`__templ_closeClothingSizeDialog_40a3`),
	}
}

func ImageCarousel(product models.Product) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var15 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var15 == nil {
			templ_7745c5c3_Var15 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 32, "<link rel=\"stylesheet\" href=\"/static/styles/user/components/carousel.css\"><div class=\"carousel-container\"><div class=\"main-image-container\"><img src=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var16 string
		templ_7745c5c3_Var16, templ_7745c5c3_Err = templ.JoinStringErrs("/static/images/" + product.Id + "-main")
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/user/Item.templ`, Line: 342, Col: 58}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var16))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 33, "\" alt=\"Main Image\" class=\"main-image\" id=\"mainImage\"></div><div class=\"preview-panel\"><button class=\"arrows prev\">&#10094;</button><div class=\"preview-container\"><img src=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var17 string
		templ_7745c5c3_Var17, templ_7745c5c3_Err = templ.JoinStringErrs("/static/images/" + product.Id + "-main")
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/user/Item.templ`, Line: 347, Col: 62}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var17))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 34, "\" alt=\"Preview 1\" class=\"preview-image active\"> <img src=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var18 string
		templ_7745c5c3_Var18, templ_7745c5c3_Err = templ.JoinStringErrs("/static/images/" + product.Id + "-1")
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/user/Item.templ`, Line: 348, Col: 59}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var18))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 35, "\" alt=\"Preview 2\" class=\"preview-image\"> <img src=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var19 string
		templ_7745c5c3_Var19, templ_7745c5c3_Err = templ.JoinStringErrs("/static/images/" + product.Id + "-2")
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/user/Item.templ`, Line: 349, Col: 59}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var19))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 36, "\" alt=\"Preview 3\" class=\"preview-image\" onerror=\"this.style.display=&#39;none&#39;;\"> <img src=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var20 string
		templ_7745c5c3_Var20, templ_7745c5c3_Err = templ.JoinStringErrs("/static/images/" + product.Id + "-3")
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/user/Item.templ`, Line: 350, Col: 59}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var20))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 37, "\" alt=\"Preview 4\" class=\"preview-image\" onerror=\"this.style.display=&#39;none&#39;;\"> <img src=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var21 string
		templ_7745c5c3_Var21, templ_7745c5c3_Err = templ.JoinStringErrs("/static/images/" + product.Id + "-4")
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/user/Item.templ`, Line: 351, Col: 59}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var21))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 38, "\" alt=\"Preview 5\" class=\"preview-image\" onerror=\"this.style.display=&#39;none&#39;;\"> <img src=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var22 string
		templ_7745c5c3_Var22, templ_7745c5c3_Err = templ.JoinStringErrs("/static/images/" + product.Id + "-5")
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/user/Item.templ`, Line: 352, Col: 59}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var22))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 39, "\" alt=\"Preview 6\" class=\"preview-image\" onerror=\"this.style.display=&#39;none&#39;;\"> <img src=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var23 string
		templ_7745c5c3_Var23, templ_7745c5c3_Err = templ.JoinStringErrs("/static/images/" + product.Id + "-6")
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/user/Item.templ`, Line: 353, Col: 59}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var23))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 40, "\" alt=\"Preview 7\" class=\"preview-image\" onerror=\"this.style.display=&#39;none&#39;;\"> <img src=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var24 string
		templ_7745c5c3_Var24, templ_7745c5c3_Err = templ.JoinStringErrs("/static/images/" + product.Id + "-7")
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/user/Item.templ`, Line: 354, Col: 59}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var24))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 41, "\" alt=\"Preview 8\" class=\"preview-image\" onerror=\"this.style.display=&#39;none&#39;;\"> <img src=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var25 string
		templ_7745c5c3_Var25, templ_7745c5c3_Err = templ.JoinStringErrs("/static/images/" + product.Id + "-8")
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/user/Item.templ`, Line: 355, Col: 59}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var25))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 42, "\" alt=\"Preview 9\" class=\"preview-image\" onerror=\"this.style.display=&#39;none&#39;;\"> <img src=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var26 string
		templ_7745c5c3_Var26, templ_7745c5c3_Err = templ.JoinStringErrs("/static/images/" + product.Id + "-9")
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/user/Item.templ`, Line: 356, Col: 59}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var26))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 43, "\" alt=\"Preview 20\" class=\"preview-image\" onerror=\"this.style.display=&#39;none&#39;;\"></div><button class=\"arrows next\">&#10095;</button></div></div><script>\n        (() => {\n            const main = document.getElementById('mainImage');\n            const container = document.querySelector('.carousel-container');\n            let previews = [...document.querySelectorAll('.preview-image')];\n            let current = 0;\n\n            const update = (index) => {\n                const visible = previews.filter(img => img.style.display !== 'none');\n                if (!visible.length) return;\n                current = ((index % visible.length) + visible.length) % visible.length;\n                main.src = visible[current].src;\n                previews.forEach((p, i) => p.classList.toggle('active', p === visible[current]));\n            };\n\n            previews.forEach(img => img.addEventListener('error', () => {\n                img.style.display = 'none';\n                update(current);\n            }));\n\n            container.addEventListener('click', e => {\n                if (e.target.classList.contains('preview-image') && e.target.style.display !== 'none')\n                    update(previews.indexOf(e.target));\n                else if (e.target.classList.contains('arrows'))\n                    update(current + (e.target.classList.contains('next') ? 1 : -1));\n            });\n\n            update(0);\n        })();\n\n    </script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
