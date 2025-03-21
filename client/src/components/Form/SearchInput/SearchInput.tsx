import { FC } from 'react';

interface Props {
    placeholder?: string;
    value?: string;
    onChange?: (value: string) => void;
}

export const SearchInput: FC<Props> = ({ placeholder, value, onChange }) => {
    return (
        <div className="search-input-container">
            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" aria-hidden="true" className="v-icon__svg">
                <path d="M9.5,3A6.5,6.5 0 0,1 16,9.5C16,11.11 15.41,12.59 14.44,13.73L14.71,14H15.5L20.5,19L19,20.5L14,15.5V14.71L13.73,14.44C12.59,15.41 11.11,16 9.5,16A6.5,6.5 0 0,1 3,9.5A6.5,6.5 0 0,1 9.5,3M9.5,5C7,5 5,7 5,9.5C5,12 7,14 9.5,14C12,14 14,12 14,9.5C14,7 12,5 9.5,5Z"></path>
            </svg>
            <input type="text" placeholder={placeholder} value={value} onChange={(e) => onChange?.(e.target.value)} />
            {value ? (
                <button className="clear" onClick={() => onChange?.('')}>
                    <svg
                        xmlns="http://www.w3.org/2000/svg"
                        viewBox="0 0 24 24"
                        aria-hidden="true"
                        className="v-icon__svg">
                        <path d="M19 6.41L17.59 5 12 10.59 6.41 5 5 6.41 10.59 12 5 17.59 6.41 19 12 13.41 17.59 19 19 17.59 13.41 12z"></path>
                    </svg>
                </button>
            ) : null}
        </div>
    );
};
