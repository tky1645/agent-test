import React, { useState, useEffect } from 'react';
import { Button } from '@/components/ui/button';
import { Plus, Loader2 } from 'lucide-react';
import { Plant } from '../types';
import { PlantCard } from './PlantCard';
import { apiService } from '../services/api';

interface PlantListProps {
  onAddPlant: () => void;
}

export const PlantList: React.FC<PlantListProps> = ({ onAddPlant }) => {
  const [plants, setPlants] = useState<Plant[]>([]);
  const [isLoading, setIsLoading] = useState(true);

  const fetchPlants = async () => {
    try {
      const plantsData = await apiService.getPlants();
      setPlants(plantsData);
    } catch (error) {
      console.error('Failed to fetch plants:', error);
    } finally {
      setIsLoading(false);
    }
  };

  useEffect(() => {
    fetchPlants();
  }, []);

  const handleWateringComplete = () => {
    fetchPlants();
  };

  if (isLoading) {
    return (
      <div className="flex items-center justify-center min-h-[400px]">
        <div className="text-center">
          <Loader2 className="h-8 w-8 animate-spin mx-auto mb-4 text-green-600" />
          <p className="text-gray-600">Loading your plants...</p>
        </div>
      </div>
    );
  }

  return (
    <div className="space-y-6">
      <div className="flex justify-between items-center">
        <div>
          <h1 className="text-3xl font-bold text-gray-900">My Plants</h1>
          <p className="text-gray-600 mt-1">
            {plants.length} {plants.length === 1 ? 'plant' : 'plants'} in your collection
          </p>
        </div>
        <Button onClick={onAddPlant} className="bg-green-600 hover:bg-green-700">
          <Plus className="mr-2 h-4 w-4" />
          Add Plant
        </Button>
      </div>

      {plants.length === 0 ? (
        <div className="text-center py-12">
          <div className="text-6xl mb-4">🌱</div>
          <h3 className="text-xl font-semibold text-gray-900 mb-2">No plants yet</h3>
          <p className="text-gray-600 mb-6">Start your plant collection by adding your first plant!</p>
          <Button onClick={onAddPlant} className="bg-green-600 hover:bg-green-700">
            <Plus className="mr-2 h-4 w-4" />
            Add Your First Plant
          </Button>
        </div>
      ) : (
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6">
          {plants.map((plant) => (
            <PlantCard
              key={plant.id}
              plant={plant}
              onWateringComplete={handleWateringComplete}
            />
          ))}
        </div>
      )}
    </div>
  );
};
