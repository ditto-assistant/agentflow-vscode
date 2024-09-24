import * as path from 'path';
import { workspace, ExtensionContext } from 'vscode';
import {
  LanguageClient,
  LanguageClientOptions,
  ServerOptions,
} from 'vscode-languageclient/node';

let client: LanguageClient;

export function activate(context: ExtensionContext) {
  const serverCommand = 'agentflow-lsp'; // Name of your Go binary
  const serverOptions: ServerOptions = {
    command: serverCommand,
    args: ['--stdio'],
  };

  const clientOptions: LanguageClientOptions = {
    documentSelector: [{ scheme: 'file', language: 'agentflow' }],
    synchronize: {
      fileEvents: workspace.createFileSystemWatcher('**/.af'),
    },
  };

  client = new LanguageClient(
    'agentflowLanguageServer',
    'AgentFlow Language Server',
    serverOptions,
    clientOptions
  );

  client.start();
}

export function deactivate(): Thenable<void> | undefined {
  if (!client) {
    return undefined;
  }
  return client.stop();
}