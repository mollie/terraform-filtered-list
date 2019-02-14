locals {
    # Include Filter
    use_include = "${length(var.include) > 0 ? 1 : 0}"
    include_input = "${distinct(var.input)}"

    # Exclude Filter
    use_exclude = "${length(var.exclude) > 0 ? 1 : 0}"
    exclude_input = "${compact(data.template_file.include.*.rendered)}"

    # Output
    filtered_list = "${compact(data.template_file.exclude.*.rendered)}"
}

data "template_file" "include" {
  # Render the template once for each item
  count    = "${length(local.include_input)}"
  template = "$${value}"
  vars {
    value = "${ !local.use_include || local.use_include && contains(var.include, local.include_input[count.index]) ? local.include_input[count.index] : ""}"
  }
}

data "template_file" "exclude" {
  # Render the template once for each item
  count    = "${length(local.exclude_input)}"
  template = "$${value}"
  vars {
    value = "${ !local.use_exclude || local.use_exclude && !contains(var.exclude, local.exclude_input[count.index]) ? local.exclude_input[count.index] : ""}"
  }
}