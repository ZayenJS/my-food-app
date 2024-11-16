export abstract class Random {
    public static buildRandomId(): string {
        const firstPart = Math.random().toString(36).substring(2, 15);
        const secondPart = Math.random().toString(36).substring(2, 15);

        return `_${firstPart}${secondPart}`;
    }
}
