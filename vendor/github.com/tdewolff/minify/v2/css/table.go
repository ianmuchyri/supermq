package css

var optionalZeroDimension = map[string]bool{
	"px":   true,
	"mm":   true,
	"q":    true,
	"cm":   true,
	"in":   true,
	"pt":   true,
	"pc":   true,
	"ch":   true,
	"em":   true,
	"ex":   true,
	"rem":  true,
	"vh":   true,
	"vw":   true,
	"vmin": true,
	"vmax": true,
	"deg":  true,
	"grad": true,
	"rad":  true,
	"turn": true,
}

// Uses http://www.w3.org/TR/2010/PR-css3-color-20101028/ for colors

// ShortenColorHex maps a color hexcode to its shorter name
var ShortenColorHex = map[string][]byte{
	"#000080": []byte("navy"),
	"#008000": []byte("green"),
	"#008080": []byte("teal"),
	"#4b0082": []byte("indigo"),
	"#800000": []byte("maroon"),
	"#800080": []byte("purple"),
	"#808000": []byte("olive"),
	"#808080": []byte("gray"),
	"#a0522d": []byte("sienna"),
	"#a52a2a": []byte("brown"),
	"#c0c0c0": []byte("silver"),
	"#cd853f": []byte("peru"),
	"#d2b48c": []byte("tan"),
	"#da70d6": []byte("orchid"),
	"#dda0dd": []byte("plum"),
	"#ee82ee": []byte("violet"),
	"#f0e68c": []byte("khaki"),
	"#f0ffff": []byte("azure"),
	"#f5deb3": []byte("wheat"),
	"#f5f5dc": []byte("beige"),
	"#fa8072": []byte("salmon"),
	"#faf0e6": []byte("linen"),
	"#ff6347": []byte("tomato"),
	"#ff7f50": []byte("coral"),
	"#ffa500": []byte("orange"),
	"#ffc0cb": []byte("pink"),
	"#ffd700": []byte("gold"),
	"#ffe4c4": []byte("bisque"),
	"#fffafa": []byte("snow"),
	"#fffff0": []byte("ivory"),
	"#ff0000": []byte("red"),
	"#f00":    []byte("red"),
}

// ShortenColorName maps a color name to its shorter hexcode
var ShortenColorName = map[Hash][]byte{
	Black:                []byte("#000"),
	Darkblue:             []byte("#00008b"),
	Mediumblue:           []byte("#0000cd"),
	Darkgreen:            []byte("#006400"),
	Darkcyan:             []byte("#008b8b"),
	Deepskyblue:          []byte("#00bfff"),
	Darkturquoise:        []byte("#00ced1"),
	Mediumspringgreen:    []byte("#00fa9a"),
	Springgreen:          []byte("#00ff7f"),
	Midnightblue:         []byte("#191970"),
	Dodgerblue:           []byte("#1e90ff"),
	Lightseagreen:        []byte("#20b2aa"),
	Forestgreen:          []byte("#228b22"),
	Seagreen:             []byte("#2e8b57"),
	Darkslategray:        []byte("#2f4f4f"),
	Limegreen:            []byte("#32cd32"),
	Mediumseagreen:       []byte("#3cb371"),
	Turquoise:            []byte("#40e0d0"),
	Royalblue:            []byte("#4169e1"),
	Steelblue:            []byte("#4682b4"),
	Darkslateblue:        []byte("#483d8b"),
	Mediumturquoise:      []byte("#48d1cc"),
	Darkolivegreen:       []byte("#556b2f"),
	Cadetblue:            []byte("#5f9ea0"),
	Cornflowerblue:       []byte("#6495ed"),
	Mediumaquamarine:     []byte("#66cdaa"),
	Slateblue:            []byte("#6a5acd"),
	Olivedrab:            []byte("#6b8e23"),
	Slategray:            []byte("#708090"),
	Lightslateblue:       []byte("#789"),
	Mediumslateblue:      []byte("#7b68ee"),
	Lawngreen:            []byte("#7cfc00"),
	Chartreuse:           []byte("#7fff00"),
	Aquamarine:           []byte("#7fffd4"),
	Lightskyblue:         []byte("#87cefa"),
	Blueviolet:           []byte("#8a2be2"),
	Darkmagenta:          []byte("#8b008b"),
	Saddlebrown:          []byte("#8b4513"),
	Darkseagreen:         []byte("#8fbc8f"),
	Lightgreen:           []byte("#90ee90"),
	Mediumpurple:         []byte("#9370db"),
	Darkviolet:           []byte("#9400d3"),
	Palegreen:            []byte("#98fb98"),
	Darkorchid:           []byte("#9932cc"),
	Yellowgreen:          []byte("#9acd32"),
	Darkgray:             []byte("#a9a9a9"),
	Lightblue:            []byte("#add8e6"),
	Greenyellow:          []byte("#adff2f"),
	Paleturquoise:        []byte("#afeeee"),
	Lightsteelblue:       []byte("#b0c4de"),
	Powderblue:           []byte("#b0e0e6"),
	Firebrick:            []byte("#b22222"),
	Darkgoldenrod:        []byte("#b8860b"),
	Mediumorchid:         []byte("#ba55d3"),
	Rosybrown:            []byte("#bc8f8f"),
	Darkkhaki:            []byte("#bdb76b"),
	Mediumvioletred:      []byte("#c71585"),
	Indianred:            []byte("#cd5c5c"),
	Chocolate:            []byte("#d2691e"),
	Lightgray:            []byte("#d3d3d3"),
	Goldenrod:            []byte("#daa520"),
	Palevioletred:        []byte("#db7093"),
	Gainsboro:            []byte("#dcdcdc"),
	Burlywood:            []byte("#deb887"),
	Lightcyan:            []byte("#e0ffff"),
	Lavender:             []byte("#e6e6fa"),
	Darksalmon:           []byte("#e9967a"),
	Palegoldenrod:        []byte("#eee8aa"),
	Lightcoral:           []byte("#f08080"),
	Aliceblue:            []byte("#f0f8ff"),
	Honeydew:             []byte("#f0fff0"),
	Sandybrown:           []byte("#f4a460"),
	Whitesmoke:           []byte("#f5f5f5"),
	Mintcream:            []byte("#f5fffa"),
	Ghostwhite:           []byte("#f8f8ff"),
	Antiquewhite:         []byte("#faebd7"),
	Lightgoldenrodyellow: []byte("#fafad2"),
	Fuchsia:              []byte("#f0f"),
	Magenta:              []byte("#f0f"),
	Deeppink:             []byte("#ff1493"),
	Orangered:            []byte("#ff4500"),
	Darkorange:           []byte("#ff8c00"),
	Lightsalmon:          []byte("#ffa07a"),
	Lightpink:            []byte("#ffb6c1"),
	Peachpuff:            []byte("#ffdab9"),
	Navajowhite:          []byte("#ffdead"),
	Moccasin:             []byte("#ffe4b5"),
	Mistyrose:            []byte("#ffe4e1"),
	Blanchedalmond:       []byte("#ffebcd"),
	Papayawhip:           []byte("#ffefd5"),
	Lavenderblush:        []byte("#fff0f5"),
	Seashell:             []byte("#fff5ee"),
	Cornsilk:             []byte("#fff8dc"),
	Lemonchiffon:         []byte("#fffacd"),
	Floralwhite:          []byte("#fffaf0"),
	Yellow:               []byte("#ff0"),
	Lightyellow:          []byte("#ffffe0"),
	White:                []byte("#fff"),
}

// PropertyOverrides is a map of which properties are overridden by the given property.
var PropertyOverrides = map[Hash][]Hash{
	Background:      {Background, Background_Image, Background_Position, Background_Size, Background_Repeat, Background_Origin, Background_Clip, Background_Attachment, Background_Color},
	Font:            {Font, Font_Style, Font_Variant, Font_Weight, Font_Stretch, Font_Size, Font_Family, Line_Height},
	Border:          {Border, Border_Width, Border_Top_Width, Border_Right_Width, Border_Bottom_Width, Border_Left_Width, Border_Style, Border_Top_Style, Border_Right_Style, Border_Bottom_Style, Border_Left_Style, Border_Color, Border_Top_Color, Border_Right_Color, Border_Bottom_Color, Border_Left_Color},
	Border_Width:    {Border_Width, Border_Top_Width, Border_Right_Width, Border_Bottom_Width, Border_Left_Width},
	Border_Style:    {Border_Style, Border_Top_Style, Border_Right_Style, Border_Bottom_Style, Border_Left_Style},
	Border_Color:    {Border_Color, Border_Top_Color, Border_Right_Color, Border_Bottom_Color, Border_Left_Color},
	Border_Top:      {Border_Top, Border_Top_Width, Border_Top_Style, Border_Top_Color},
	Border_Right:    {Border_Right, Border_Right_Width, Border_Right_Style, Border_Right_Color},
	Border_Bottom:   {Border_Bottom, Border_Bottom_Width, Border_Bottom_Style, Border_Bottom_Color},
	Border_Left:     {Border_Left, Border_Left_Width, Border_Left_Style, Border_Left_Color},
	Margin:          {Margin, Margin_Top, Margin_Right, Margin_Bottom, Margin_Left},
	Padding:         {Padding, Padding_Top, Padding_Right, Padding_Bottom, Padding_Left},
	Column_Rule:     {Column_Rule, Column_Rule_Width, Column_Rule_Style, Column_Rule_Color},
	Animation:       {Animation, Animation_Name, Animation_Duration, Animation_Timing_Function, Animation_Delay, Animation_Iteration_Count, Animation_Direction, Animation_Fill_Mode, Animation_Play_State},
	Columns:         {Columns, Column_Width, Column_Count},
	Flex:            {Flex, Flex_Basis, Flex_Grow, Flex_Shrink},
	Flex_Flow:       {Flex_Flow, Flex_Direction, Flex_Wrap},
	Grid:            {Grid, Grid_Template_Rows, Grid_Template_Columns, Grid_Template_Areas, Grid_Auto_Rows, Grid_Auto_Columns, Grid_Auto_Flow, Grid_Column_Gap, Grid_Row_Gap, Column_Gap, Row_Gap},
	Grid_Area:       {Grid_Area, Grid_Row_Start, Grid_Column_Start, Grid_Row_End, Grid_Column_End},
	Grid_Row:        {Grid_Row, Grid_Row_Start, Grid_Row_End},
	Grid_Column:     {Grid_Column, Grid_Column_Start, Grid_Column_End},
	Grid_Template:   {Grid_Template, Grid_Template_Rows, Grid_Template_Columns, Grid_Template_Areas},
	List_Style:      {List_Style, List_Style_Image, List_Style_Position, List_Style_Type},
	Offset:          {Offset, Offset_Position, Offset_Path, Offset_Distance, Offset_Anchor, Offset_Rotate},
	Outline:         {Outline, Outline_Width, Outline_Style, Outline_Color},
	Overflow:        {Overflow, Overflow_X, Overflow_Y},
	Place_Content:   {Place_Content, Align_Content, Justify_Content},
	Place_Items:     {Place_Items, Align_Items, Justify_Items},
	Place_Self:      {Place_Self, Align_Self, Justify_Self},
	Text_Decoration: {Text_Decoration, Text_Decoration_Color, Text_Decoration_Color, Text_Decoration_Line, Text_Decoration_Thickness},
	Transition:      {Transition, Transition_Property, Transition_Duration, Transition_Timing_Function, Transition_Delay},
}