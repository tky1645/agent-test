import { Plant, PlantCreate, WateringRecord, WateringRecordCreate, PlantStatus } from '../types';

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080/api/v1';

class ApiService {

  async getPlants(): Promise<Plant[]> {
    const response = await fetch(`${API_BASE_URL}/plants`);
    if (!response.ok) {
      throw new Error('Failed to fetch plants');
    }
    return response.json();
  }

  async createPlant(plantData: PlantCreate): Promise<Plant> {
    const response = await fetch(`${API_BASE_URL}/plants`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(plantData),
    });
    if (!response.ok) {
      throw new Error('Failed to create plant');
    }
    return response.json();
  }

  async getPlantStatus(plantId: string): Promise<PlantStatus> {
    const response = await fetch(`${API_BASE_URL}/plants/${plantId}`);
    if (!response.ok) {
      throw new Error('Plant not found');
    }
    const plant = await response.json();
    
    const wateringResponse = await fetch(`${API_BASE_URL}/plants/${plantId}/watering`);
    const wateringRecords = wateringResponse.ok ? await wateringResponse.json() : [];
    
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
    const response = await fetch(`${API_BASE_URL}/plants/${plantId}/watering`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(recordData),
    });
    if (!response.ok) {
      throw new Error('Failed to add watering record');
    }
    return response.json();
  }

  async getWateringRecords(plantId: string): Promise<WateringRecord[]> {
    const response = await fetch(`${API_BASE_URL}/plants/${plantId}/watering`);
    if (!response.ok) {
      throw new Error('Failed to fetch watering records');
    }
    return response.json();
  }
}

export const apiService = new ApiService();
