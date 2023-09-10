/* --------------------------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Licensed under the MIT License. See License.txt in the project root for license information.
 * ------------------------------------------------------------------------------------------ */

import * as path from "path";
import { workspace, ExtensionContext } from "vscode";
import * as vscode from "vscode";

import {
  Executable,
  LanguageClient,
  LanguageClientOptions,
  ServerOptions,
  TransportKind,
} from "vscode-languageclient/node";

let client: LanguageClient;

export function activate(context: ExtensionContext) {
  const goBin = context.asAbsolutePath(path.join("..", "bin", "langserver"));
  const cwd = context.asAbsolutePath("..");

  console.log(`Starting ${goBin}`);

  const options: Executable = {
    command: goBin,
    options: {
      cwd: cwd,
    },
  };
  // If the extension is launched in debug mode then the debug server options are used
  // Otherwise the run options are used
  const serverOptions: ServerOptions = {
    run: options,
    debug: options,
  };

  // Options to control the language client
  const clientOptions: LanguageClientOptions = {
    // Register the server for hcl documents
    documentSelector: [{ scheme: "file", language: "hcl" }],
    synchronize: {
      // Notify the server about file changes to '.clientrc files contained in the workspace
      fileEvents: workspace.createFileSystemWatcher("**/.clientrc"),
    },
  };

  // Create the language client and start the client.
  client = new LanguageClient(
    "terragrunt",
    "Terragrunt",
    serverOptions,
    clientOptions
  );

  // Start the client. This will also launch the server
  client.start();
}

export function deactivate(): Thenable<void> | undefined {
  if (!client) {
    return undefined;
  }
  return client.stop();
}
