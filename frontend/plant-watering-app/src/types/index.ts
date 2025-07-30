export interface User {
  id: string;
  cognito_id: string;
  email: string;
  name: string;
  created_at: string;
  updated_at: string;
}

export interface Plant {
  id: string;
  user_id: string;
  name: string;
  description?: string;
  image_url?: string;
  created_at: string;
  updated_at: string;
}

export interface PlantCreate {
  name: string;
  description?: string;
  image_url?: string;
}

export interface WateringRecord {
  id: string;
  plant_id: string;
  watered_at: string;
  notes?: string;
  created_at: string;
}

export interface WateringRecordCreate {
  watered_at: string;
  notes?: string;
}

export interface PlantStatus {
  id: string;
  name: string;
  last_watered_at?: string;
  days_since_last_watering?: number | null;
  previous_watering_interval?: number;
}

export interface AuthUser {
  id: string;
  email: string;
  name: string;
  token: string;
}

export interface ApiResponse<T> {
  data: T;
  success: boolean;
  message?: string;
}
