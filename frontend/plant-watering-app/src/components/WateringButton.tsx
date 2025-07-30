import React, { useState } from 'react';
import { Button } from '@/components/ui/button';
import { Droplets, Loader2, Check } from 'lucide-react';
import { apiService } from '../services/api';

interface WateringButtonProps {
  plantId: string;
  onWateringComplete: () => void;
  disabled?: boolean;
}

export const WateringButton: React.FC<WateringButtonProps> = ({
  plantId,
  onWateringComplete,
  disabled = false
}) => {
  const [isWatering, setIsWatering] = useState(false);
  const [justWatered, setJustWatered] = useState(false);

  const handleWatering = async () => {
    setIsWatering(true);
    try {
      await apiService.addWateringRecord(plantId, {
        watered_at: new Date().toISOString(),
        notes: 'Quick watering'
      });
      
      setJustWatered(true);
      onWateringComplete();
      
      setTimeout(() => {
        setJustWatered(false);
      }, 2000);
    } catch (error) {
      console.error('Failed to record watering:', error);
    } finally {
      setIsWatering(false);
    }
  };

  if (justWatered) {
    return (
      <Button variant="outline" className="text-green-600 border-green-200 bg-green-50" disabled>
        <Check className="mr-2 h-4 w-4" />
        Watered!
      </Button>
    );
  }

  return (
    <Button 
      onClick={handleWatering} 
      disabled={disabled || isWatering}
      className="bg-blue-600 hover:bg-blue-700"
    >
      {isWatering ? (
        <>
          <Loader2 className="mr-2 h-4 w-4 animate-spin" />
          Watering...
        </>
      ) : (
        <>
          <Droplets className="mr-2 h-4 w-4" />
          Water Now
        </>
      )}
    </Button>
  );
};
