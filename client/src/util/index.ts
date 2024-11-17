export const buildClassName = (...classNames: (string | undefined)[]): string => {
    return classNames.join(' ').trim();
};
