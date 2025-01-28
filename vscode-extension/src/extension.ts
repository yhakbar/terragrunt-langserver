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
  let goBin: string, cwd: string;

  const isDevelopmentMode =
    context.extensionMode === vscode.ExtensionMode.Development;
  if (isDevelopmentMode) {
    // Run the langserver wrapper script in the parent directory. This will re-compile the langserver
    // every time the extension is started
    goBin = context.asAbsolutePath(path.join("..", "bin", "langserver"));
    cwd = context.asAbsolutePath("..");
  } else {
    // Run the langserver binary that the build script copies into the extension's bin directory.
    goBin = context.asAbsolutePath(path.join("out", "langserver"));
    cwd = context.asAbsolutePath(".");
  }

  const options: Executable = {
    command: goBin,
    options: {
      cwd: cwd,
    },
    transport: TransportKind.stdio,
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
    documentSelector: [{ scheme: "file", pattern: "**/*.hcl" }],
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

  const log = (message: string) => {
    client.outputChannel.appendLine(`client: ${message}`);
  };
  log(`Starting ${goBin}`);

  // Start the client. This will also launch the server
  client.start();
}

export function deactivate(): Thenable<void> | undefined {
  if (!client) {
    return undefined;
  }
  return client.stop();
}
