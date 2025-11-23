export function idTimeFormat(dateStr: string) {
  const dateFormatter = new Intl.DateTimeFormat("id-ID", {
    weekday: "long",
    year: "numeric",
    month: "long",
    day: "numeric",
  });
  return dateFormatter.format(new Date(dateStr));
}
