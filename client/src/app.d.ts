// app.d.ts
declare global {
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
}

export {};
