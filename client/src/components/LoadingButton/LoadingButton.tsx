import { ForwardedRef, forwardRef } from 'react';

interface Props {
    children: React.ReactNode;
    classNames?: {
        root?: string;
        loading?: string;
    };
    type: 'button' | 'submit' | 'reset';
    loading: boolean;
    onClick?: () => void;
}

export const LoadingButton = forwardRef(
    ({ children, classNames, type, loading, onClick }: Props, ref: ForwardedRef<HTMLButtonElement>) => {
        return (
            <button ref={ref} className={classNames?.root} type={type} disabled={loading} onClick={onClick}>
                {loading ? <span className={classNames?.loading} /> : children}
            </button>
        );
    },
);
