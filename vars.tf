variable "input" {
    type = "list"
    description = "The input list"
}

variable "include" {
    type = "list"
    default = []
    description = "The include list, returns all items from input that are also in this list"
}

variable "exclude" {
    type = "list"
    default = []
    description = "The exclude list, returns all items from input excluding this list"
}