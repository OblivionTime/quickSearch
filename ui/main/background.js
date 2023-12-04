const {
  app,
  BrowserWindow,
  Tray,
  Menu,
  nativeImage,
  globalShortcut,
} = require("electron");
const { join, dirname } = require("path");
const { spawn } = require("child_process");
let tray;
let exampleProcess;
// 屏蔽安全警告
// ectron Security Warning (Insecure Content-Security-Policy)
process.env["ELECTRON_DISABLE_SECURITY_WARNINGS"] = "true";
const createWindow = () => {
  const icon = nativeImage.createFromPath(join(__dirname, "logo.png"));
  const win = new BrowserWindow({
    width: 1200,
    useContentSize: true,
    height: 600,
    icon: icon,
    autoHideMenuBar: true,
    minWidth: 1200,
    minHeight: 600,
    webPreferences: {
      webSecurity: false,
      nodeIntegration: true,
      enableRemoteModule: true,
      contextIsolation: false,
    },
  });
  globalShortcut.register("Control+Alt+Space", () => {
    win.show();
  });
  tray = new Tray(icon);

  const contextMenu = Menu.buildFromTemplate([
    { label: "主界面", click: () => win.show() },
    { label: "退出", click: () => app.exit() },
  ]);
  tray.on("double-click", () => {
    win.show();
  });
  tray.setToolTip("快寻,快速寻找你想要的");
  tray.setContextMenu(contextMenu);
  win.on("close", (event) => {
    win.hide();
    win.setSkipTaskbar(true);
    event.preventDefault();
  });
  // development模式
  if (process.env.VITE_DEV_SERVER_URL) {
    win.loadURL(process.env.VITE_DEV_SERVER_URL);
    // 开启调试台
    // win.webContents.openDevTools()
  } else {
    const appPath = app.getAppPath();
    process.chdir(appPath + "/../../");

    exampleProcess = spawn("quickSearch.exe");
    // win.setMenu(null);
    win.loadFile(join(__dirname, "../dist/index.html"));
  }
};
app.on(
  "certificate-error",
  (event, webContents, url, error, certificate, callback) => {
    log("certificate-error");
    //允许私有证书
    event.preventDefault();
    callback(true);
  }
);
// 忽略证书相关错误
app.commandLine.appendSwitch("ignore-certificate-errors");
app.whenReady().then(() => {
  createWindow();
  app.on("activate", () => {
    if (BrowserWindow.getAllWindows().length === 0) createWindow();
  });
});
app.on("window-all-closed", () => {
  try {
    if (exampleProcess) {
      process.kill(exampleProcess.pid);
    }
  } catch (error) {}
  if (process.platform !== "darwin") app.quit();
});
