locals {
    # Intersect Filter
    use_intersect = "${length(var.intersect) > 0 && var.intersect[0] != "" ? true : false}"
    intersect_input = "${distinct(var.input)}"

    # Exclude Filter
    use_exclude = "${length(var.exclude) > 0 && var.exclude[0] != "" ? true : false}"
    exclude_input = "${compact(data.template_file.intersect.*.rendered)}"

    # Output
    filtered_list = "${compact(data.template_file.exclude.*.rendered)}"
}

data "template_file" "intersect" {
  # Render the template once for each item
  count    = "${length(local.intersect_input)}"
  template = "$${value}"
  vars {
    value = "${ !local.use_intersect || local.use_intersect && contains(var.intersect, local.intersect_input[count.index]) ? local.intersect_input[count.index] : ""}"
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