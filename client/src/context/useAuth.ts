import { AuthCtx } from './authProvider';
import { useContext } from 'react';

export const useAuth = () => useContext(AuthCtx);
