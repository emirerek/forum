function getUserLocale(): string | null {
    if (typeof window === "undefined") return null;
    if (navigator.languages != undefined)
        return navigator.languages[0];
    return navigator.language;
}

function formatDate(iso8601: string): string {
    const locale = getUserLocale();
    return new Date(iso8601).toLocaleDateString(locale || "en-US");
}

export { getUserLocale, formatDate };