import axios from 'axios';

const API_URL = import.meta.env.VITE_API_URL || 'http://localhost:8000';

const apiClient = axios.create({
    baseURL: API_URL,
    headers: {
        'Content-Type': 'application/json'
    }
});

// Add request interceptor to include auth token
apiClient.interceptors.request.use((config) => {
    const token = localStorage.getItem('sb-auth-token');
    if (token) {
        config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
});

export const api = {
    // Posts
    getPosts: (params?: { category?: string, published_only?: boolean, author_id?: string, limit?: number, offset?: number }) =>
        apiClient.get('/posts', { params }),
    
    getPost: (slug: string) => 
        apiClient.get(`/posts/${slug}`),
    
    createPost: (data: any) =>
        apiClient.post('/posts', data),
    
    updatePost: (id: string, data: any) =>
        apiClient.put(`/posts/${id}`, data),
    
    deletePost: (id: string) =>
        apiClient.delete(`/posts/${id}`),

    // Comments
    getComments: (postId: string) =>
        apiClient.get(`/posts/${postId}/comments`),
    
    createComment: (postId: string, data: any) =>
        apiClient.post(`/posts/${postId}/comments`, data),

    // Categories
    getCategories: () =>
        apiClient.get('/categories'),
    
    createCategory: (data: any) =>
        apiClient.post('/categories', data),

    // Profiles
    getProfile: (userId: string) =>
        apiClient.get(`/profiles/${userId}`),
    
    updateProfile: (userId: string, data: any) =>
        apiClient.put(`/profiles/${userId}`, data),

    // Saved Posts
    savePost: (postId: string) =>
        apiClient.post(`/posts/${postId}/save`),
    
    unsavePost: (postId: string) =>
        apiClient.delete(`/posts/${postId}/unsave`),
    
    getSavedPosts: () =>
        apiClient.get('/saved-posts'),

    // Website Content
    getWebsiteContent: (pageName: string) =>
        apiClient.get(`/website-content/${pageName}`),
    
    createWebsiteContent: (data: any) =>
        apiClient.post('/website-content', data),
};
