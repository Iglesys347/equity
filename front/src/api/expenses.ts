import instance from "./axios_instance"

export function getExpenses() {
    return instance.get("/expenses")
}