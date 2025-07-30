export class S3Service {
  static async uploadImage(file: File, userId: string): Promise<string> {
    await new Promise(resolve => setTimeout(resolve, 2000));
    
    console.log(`Mock: Uploading ${file.name} for user ${userId}`);
    
    const mockImageUrls = [
      'https://images.unsplash.com/photo-1416879595882-3373a0480b5b?w=400',
      'https://images.unsplash.com/photo-1485955900006-10f4d324d411?w=400',
      'https://images.unsplash.com/photo-1501004318641-b39e6451bec6?w=400',
      'https://images.unsplash.com/photo-1463320726281-696a485928c7?w=400',
      'https://images.unsplash.com/photo-1558618666-fcd25c85cd64?w=400'
    ];
    
    return mockImageUrls[Math.floor(Math.random() * mockImageUrls.length)];
  }

  static async deleteImage(imageUrl: string): Promise<void> {
    await new Promise(resolve => setTimeout(resolve, 500));
    console.log('Mock: Deleted image', imageUrl);
  }
}
