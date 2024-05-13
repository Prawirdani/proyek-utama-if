import { AuthCtx } from '@/providers/authProvider';
import { useContext } from 'react';

export const useAuth = () => useContext(AuthCtx);
