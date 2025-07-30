import React, { useState } from 'react';
import { AuthProvider, useAuth } from './contexts/AuthContext';
import { LoginForm } from './components/LoginForm';
import { Header } from './components/Header';
import { PlantList } from './components/PlantList';
import { PlantRegistration } from './components/PlantRegistration';
import { Toaster } from '@/components/ui/toaster';

type View = 'plants' | 'add-plant';

const AppContent: React.FC = () => {
  const { user, isLoading } = useAuth();
  const [currentView, setCurrentView] = useState<View>('plants');

  if (isLoading) {
    return (
      <div className="min-h-screen flex items-center justify-center">
        <div className="text-center">
          <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-green-600 mx-auto mb-4"></div>
          <p className="text-gray-600">Loading...</p>
        </div>
      </div>
    );
  }

  if (!user) {
    return <LoginForm />;
  }

  const handleAddPlant = () => {
    setCurrentView('add-plant');
  };

  const handleBackToPlants = () => {
    setCurrentView('plants');
  };

  const handlePlantAdded = () => {
    setCurrentView('plants');
  };

  return (
    <div className="min-h-screen bg-gray-50">
      <Header />
      <main className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        {currentView === 'plants' && (
          <PlantList onAddPlant={handleAddPlant} />
        )}
        {currentView === 'add-plant' && (
          <PlantRegistration
            onBack={handleBackToPlants}
            onPlantAdded={handlePlantAdded}
          />
        )}
      </main>
      <Toaster />
    </div>
  );
};

function App() {
  return (
    <AuthProvider>
      <AppContent />
    </AuthProvider>
  );
}

export default App;
