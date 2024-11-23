import React, { useState } from 'react';
import { useZxing } from 'react-zxing';

interface BarcodeScannerProps {
    onDetected: (code: string) => void;
}

export const BarcodeScanner: React.FC<BarcodeScannerProps> = ({ onDetected }) => {
    const [isScanning, setIsScanning] = useState(true);
    const [error, setError] = useState<string | null>(null);

    // Set up the `useZxing` hook with error handling
    const { ref } = useZxing({
        onResult(result) {
            const code = result.getText();
            setIsScanning(false); // Stop scanning after detection
            onDetected(code);
        },
        onError(err) {
            console.error('Error scanning barcode:', err);
            setError('Failed to access the camera or scan barcode. Please check permissions.');
        },
    });

    return (
        <div style={{ width: '100%', maxWidth: 500, margin: '0 auto', textAlign: 'center' }}>
            {isScanning ? (
                <video ref={ref} style={{ width: '100%' }} />
            ) : (
                <button onClick={() => setIsScanning(true)}>Scan Again</button>
            )}
            {error && <p style={{ color: 'red' }}>{error}</p>}
        </div>
    );
};
