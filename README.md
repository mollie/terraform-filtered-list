# Terraform filtered list
Let's say we have three lists:
```
locals {
    input = [ "a", "b", "c", "d", "e" ]
    include = [ "d", "e", "f" ]
    exclude = [ "a", "b", "e" ]
}
```

## Include
And we want all items that are both in `input` and in `include`
```
module "selection" {
  source = "git::ssh://git@github.com/TheWolfNL/terraform-filtered-list.git"
  input = "${local.input}"
  include = "${local.include}"
}
```
Then the output would be:
```
    module.selection.filtered_list = [
        d,
        e
    ]
```

## Exclude
And now we want all items that are in `input` and **NOT** in `exclude`
```
module "selection" {
  source = "git::ssh://git@github.com/TheWolfNL/terraform-filtered-list.git"
  input = "${local.input}"
  exclude = "${local.exclude}"
}
```
Then the output would be:
```
    module.selection.filtered_list = [
        c,
        d
    ]
```

## Include and Exclude
Now we'd like all items that are both in `input` and in `include` and **NOT** in `exclude`
```
module "selection" {
  source = "git::ssh://git@github.com/TheWolfNL/terraform-filtered-list.git"
  input = "${local.input}"
  include = "${local.include}"
  exclude = "${local.exclude}"
}
```
Then the output would be:
```
    module.selection.filtered_list = [
        d
    ]
```

## Output
The output will be available as `filtered_list`.
```
"${module.selection.filtered_list}"
```

## Usage
If you encounter the following error:
> should be a list
You'll need to add `[]` around the output, for terraform to work properly.
```
list = ["${module.selection.filtered_list}"]
```