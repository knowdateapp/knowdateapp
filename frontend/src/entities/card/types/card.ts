import { Workspace } from 'shared/types';

export type Card = {
  id: string;
  answer: string;
  question: string;
  workspace: Workspace;
};
