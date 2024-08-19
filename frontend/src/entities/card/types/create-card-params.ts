import { Workspace } from 'shared/types';

export type CreateCardParams = {
  workspace: Workspace;
  question: string;
  answer: string;
};
