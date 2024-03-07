/// <reference types="vite/client" />
/* eslint-disable @typescript-eslint/no-empty-interface */

import '@emotion/react';
import type { GlobalToken } from 'antd/es/theme/interface';

declare module '@emotion/react' {
  export interface Theme extends GlobalToken {}
}
