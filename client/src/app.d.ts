// app.d.ts
declare global {
  type ApiResponse<T> = {
    data: T | null;
    message?: string;
  };

  type Pagination = {
    page: number;
    size: number;
    totalData: number;
    maxPage: number;
  };

  type ErrorResponse = {
    error: {
      code: number;
      message: string;
      details?: Record<string, string>;
    };
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

  type TipePembayaran = 'TUNAI' | 'CARD' | 'MOBILE';
  type MetodePembayaran = {
    id: number;
    tipePembayaran: TipePembayaran;
    metode: string;
    deskripsi: string;
  };

  type User = {
    id: number;
    nama: string;
    username: string;
    active: boolean;
    role: string;
    createdAt: Date;
    updatedAt: Date;
  };

  type DetailTransaksi = {
    id: number;
    namaMenu: string;
    hargaMenu: number;
    kuantitas: number;
    subtotal: number;
  };

  type statusTransaksi = 'Diproses' | 'Selesai' | 'Batal';
  type Transaksi = {
    id: number;
    namaPelanggan: string;
    kasir: string;
    meja: Meja;
    tipe: string;
    status: statusTransaksi;
    catatan: string;
    detail: DetailTransaksi[];
    total: number;
    waktuPesanan: Date;
  };
}

export {};
