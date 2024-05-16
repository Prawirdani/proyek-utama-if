// app.d.ts
declare global {
  type ApiResponse<T> = {
    data: T | null;
    message?: string;
  };

  type ErrorResponse = {
    code: number;
    message: string;
    details?: Record<string, string>;
  };

  type UserRole = 'Manajer' | 'Kasir';
  type User = {
    id: number;
    nama: string;
    username: string;
    password: string;
    active: boolean;
    role: UserRole;
    createdAt: Date;
    updatedAt: Date;
  };

  type AuthUser = {
    id: number;
    nama: string;
    username: string;
    role: UserRole;
  };

  type Kategori = {
    id: number;
    nama: string;
    createdAt: Date;
    updatedAt: Date;
  };

  type Menu = {
    id: number;
    nama: string;
    deskripsi: string;
    harga: number;
    kategori: Kategori;
    url: string;
    createdAt: Date;
    updatedAt: Date;
  };

  type StatusMeja = 'Kosong' | 'Tersedia';
  type Meja = {
    id: number;
    nomor: string;
    status: StatusMeja;
  };
}

export {};
