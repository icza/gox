/*

Package i18n contains internationalization and localization utilities.

Translations are handled by the Dict type and its Get method, whose type is captured
by Translator. Dict presents a very (memory) compact and efficient translation
model, while preserving the flexibility of a map with one tiny compromise:
if something translates into the empty string, the Empty constant must be used
(because generally the empty string in a Dict denotes a missing translation).

Locales must be modeled with integers, where 0 denotes your default, fallback locale.

Generally Dict values are not needed to be retained, only their Get method value:

    var Day = Dict{EN: "Day", DE: "Tage"}.Get

Where Day is of type Translator, and can be called like this:

    fmt.Println("Day in English:", Day(EN))

*/
package i18n
