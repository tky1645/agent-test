import { AuthUser } from '../types';

const MOCK_USER: AuthUser = {
  id: 'user-123',
  email: 'user@example.com',
  name: 'Plant Lover',
  token: 'mock-jwt-token-123'
};

export class AuthService {
  private static readonly TOKEN_KEY = 'plant_watering_token';
  private static readonly USER_KEY = 'plant_watering_user';

  static async login(email: string, password: string): Promise<AuthUser> {
    await new Promise(resolve => setTimeout(resolve, 1000));
    
    if (email && password) {
      const user = { ...MOCK_USER, email };
      localStorage.setItem(this.TOKEN_KEY, user.token);
      localStorage.setItem(this.USER_KEY, JSON.stringify(user));
      return user;
    }
    
    throw new Error('Invalid credentials');
  }

  static async loginWithOAuth(): Promise<AuthUser> {
    await new Promise(resolve => setTimeout(resolve, 1500));
    
    const user = MOCK_USER;
    localStorage.setItem(this.TOKEN_KEY, user.token);
    localStorage.setItem(this.USER_KEY, JSON.stringify(user));
    return user;
  }

  static logout(): void {
    localStorage.removeItem(this.TOKEN_KEY);
    localStorage.removeItem(this.USER_KEY);
  }

  static getCurrentUser(): AuthUser | null {
    const userStr = localStorage.getItem(this.USER_KEY);
    return userStr ? JSON.parse(userStr) : null;
  }

  static getToken(): string | null {
    return localStorage.getItem(this.TOKEN_KEY);
  }

  static isAuthenticated(): boolean {
    return !!this.getToken();
  }
}
