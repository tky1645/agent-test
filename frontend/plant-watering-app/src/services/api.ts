import { Plant, PlantCreate, WateringRecord, WateringRecordCreate, PlantStatus } from '../types';

class ApiService {

  async getPlants(): Promise<Plant[]> {
    const mockPlants: Plant[] = JSON.parse(localStorage.getItem('mock_plants') || '[]');
    
    if (mockPlants.length === 0) {
      const defaultPlants: Plant[] = [
        {
          id: '1',
          user_id: 'user-123',
          name: 'Monstera Deliciosa',
          description: 'Beautiful tropical plant',
          image_url: 'https://images.unsplash.com/photo-1416879595882-3373a0480b5b?w=400',
          created_at: new Date(Date.now() - 7 * 24 * 60 * 60 * 1000).toISOString(),
          updated_at: new Date().toISOString()
        },
        {
          id: '2',
          user_id: 'user-123',
          name: 'Snake Plant',
          description: 'Low maintenance succulent',
          image_url: 'https://images.unsplash.com/photo-1485955900006-10f4d324d411?w=400',
          created_at: new Date(Date.now() - 14 * 24 * 60 * 60 * 1000).toISOString(),
          updated_at: new Date().toISOString()
        }
      ];
      localStorage.setItem('mock_plants', JSON.stringify(defaultPlants));
      return defaultPlants;
    }
    
    await new Promise(resolve => setTimeout(resolve, 500));
    return mockPlants;
  }

  async createPlant(plantData: PlantCreate): Promise<Plant> {
    await new Promise(resolve => setTimeout(resolve, 800));
    
    const newPlant: Plant = {
      id: Date.now().toString(),
      user_id: 'user-123',
      ...plantData,
      created_at: new Date().toISOString(),
      updated_at: new Date().toISOString()
    };

    const plants = JSON.parse(localStorage.getItem('mock_plants') || '[]');
    plants.push(newPlant);
    localStorage.setItem('mock_plants', JSON.stringify(plants));
    
    return newPlant;
  }

  async getPlantStatus(plantId: string): Promise<PlantStatus> {
    await new Promise(resolve => setTimeout(resolve, 300));
    
    const wateringRecords = JSON.parse(localStorage.getItem(`watering_${plantId}`) || '[]');
    const plants = JSON.parse(localStorage.getItem('mock_plants') || '[]');
    const plant = plants.find((p: Plant) => p.id === plantId);
    
    if (!plant) {
      throw new Error('Plant not found');
    }

    const lastRecord = wateringRecords.sort((a: WateringRecord, b: WateringRecord) => 
      new Date(b.watered_at).getTime() - new Date(a.watered_at).getTime()
    )[0];

    const lastWateredAt = lastRecord?.watered_at;
    const daysSinceLastWatering = lastWateredAt 
      ? Math.floor((Date.now() - new Date(lastWateredAt).getTime()) / (1000 * 60 * 60 * 24))
      : null;

    return {
      id: plant.id,
      name: plant.name,
      last_watered_at: lastWateredAt,
      days_since_last_watering: daysSinceLastWatering,
      previous_watering_interval: 7
    };
  }

  async addWateringRecord(plantId: string, recordData: WateringRecordCreate): Promise<WateringRecord> {
    const response = await fetch(`http://localhost:8080/plants/${plantId}/watering`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        notes: recordData.notes
      })
    });

    if (!response.ok) {
      throw new Error(`Failed to record watering: ${response.statusText}`);
    }

    const newRecord = await response.json();
    return newRecord;
  }

  async getWateringRecords(plantId: string): Promise<WateringRecord[]> {
    const response = await fetch(`http://localhost:8080/plants/${plantId}/watering`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
      }
    });

    if (!response.ok) {
      throw new Error(`Failed to get watering records: ${response.statusText}`);
    }

    const records = await response.json();
    return records;
  }
}

export const apiService = new ApiService();
