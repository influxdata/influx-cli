
# Add quotes to a string and escape any internal quotes.
# $(call with-quotes,alice bob) -> "alice bob"
# $(call with-quotes,) -> ""
# $(call with-quotes,Bobby "Drop Tables") -> "Bobby \"Drop Tables\""
define with-quotes
$(if $(1),"$(subst ",\",$(1))")
endef

# Generate a command line option in the form $(1)"$(2)", but only if $(2)
# is not empty.
# $(call with-param,-t=,) ->
# $(call with-param,-t=,Ralph Wiggum) -> -t="Ralph Wiggum"
# $(call with-param,-t=,Bobby "Drop Tables") -> -t="Bobby \"Drop Tables\""
define with-param
$(if $(2),$(1)$(call with-quotes,$(2)))
endef
