# Terraform filtered list
Let's say we have three lists:
```
locals {
    input = [ "a", "b", "c", "d", "e" ]
    intersect = [ "d", "e", "f" ]
    exclude = [ "a", "b", "e" ]
}
```

## Intersect
And we want all items that are both in `input` and in `intersect`
```
module "selection" {
  source = "git::ssh://git@github.com/mollie/terraform-filtered-list.git"
  input = "${local.input}"
  intersect = "${local.intersect}"
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
  source = "git::ssh://git@github.com/mollie/terraform-filtered-list.git"
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

## Intersect and Exclude
Now we'd like all items that are both in `input` and in `intersect` and **NOT** in `exclude`
```
module "selection" {
  source = "git::ssh://git@github.com/mollie/terraform-filtered-list.git"
  input = "${local.input}"
  intersect = "${local.intersect}"
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

## Tests
To run the tests make sure this dir is under the $GOPATH (or create a symlink)
- Run `make dev-install` to have the dependencies installed if not present.
- Run `make tests` to trigger tests.
