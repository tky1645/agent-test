import React, { useState, useEffect } from 'react';
import { Card, CardContent, CardFooter } from '@/components/ui/card';
import { Badge } from '@/components/ui/badge';
import { Calendar, Clock } from 'lucide-react';
import { Plant, PlantStatus } from '../types';
import { WateringButton } from './WateringButton';
import { apiService } from '../services/api';

interface PlantCardProps {
  plant: Plant;
  onWateringComplete: () => void;
}

export const PlantCard: React.FC<PlantCardProps> = ({ plant, onWateringComplete }) => {
  const [status, setStatus] = useState<PlantStatus | null>(null);
  const [isLoading, setIsLoading] = useState(true);

  const fetchStatus = async () => {
    try {
      const plantStatus = await apiService.getPlantStatus(plant.id);
      setStatus(plantStatus);
    } catch (error) {
      console.error('Failed to fetch plant status:', error);
    } finally {
      setIsLoading(false);
    }
  };

  useEffect(() => {
    fetchStatus();
  }, [plant.id]);

  const handleWateringComplete = () => {
    fetchStatus();
    onWateringComplete();
  };

  const getStatusColor = (days: number | null | undefined) => {
    if (days === null || days === undefined) return 'bg-gray-100 text-gray-800';
    if (days === 0) return 'bg-green-100 text-green-800';
    if (days <= 3) return 'bg-yellow-100 text-yellow-800';
    if (days <= 7) return 'bg-orange-100 text-orange-800';
    return 'bg-red-100 text-red-800';
  };

  const getStatusText = (days: number | null | undefined) => {
    if (days === null || days === undefined) return 'Never watered';
    if (days === 0) return 'Watered today';
    if (days === 1) return '1 day ago';
    return `${days} days ago`;
  };

  return (
    <Card className="overflow-hidden hover:shadow-lg transition-shadow">
      <div className="aspect-square relative overflow-hidden">
        {plant.image_url ? (
          <img
            src={plant.image_url}
            alt={plant.name}
            className="w-full h-full object-cover"
          />
        ) : (
          <div className="w-full h-full bg-gradient-to-br from-green-100 to-green-200 flex items-center justify-center">
            <span className="text-green-600 text-4xl">🌱</span>
          </div>
        )}
      </div>
      
      <CardContent className="p-4">
        <h3 className="font-semibold text-lg mb-2">{plant.name}</h3>
        {plant.description && (
          <p className="text-sm text-gray-600 mb-3">{plant.description}</p>
        )}
        
        <div className="space-y-2">
          <div className="flex items-center gap-2">
            <Calendar className="h-4 w-4 text-gray-500" />
            <span className="text-sm text-gray-600">
              Added {new Date(plant.created_at).toLocaleDateString()}
            </span>
          </div>
          
          {!isLoading && status && (
            <div className="flex items-center gap-2">
              <Clock className="h-4 w-4 text-gray-500" />
              <Badge className={getStatusColor(status.days_since_last_watering)}>
                {getStatusText(status.days_since_last_watering)}
              </Badge>
            </div>
          )}
        </div>
      </CardContent>
      
      <CardFooter className="p-4 pt-0">
        <WateringButton
          plantId={plant.id}
          onWateringComplete={handleWateringComplete}
          disabled={isLoading}
        />
      </CardFooter>
    </Card>
  );
};
