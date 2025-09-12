package exports

import "fmt"

// Return native Node.js import
func NodeImportation() string {
	return `// Node.js importations
import * as http from "http";
import * as url from "url";`
}

func NodeRouteExemple(lowerName string, upperName string) string {
	return fmt.Sprintf(`// Node.js Routes
/**
 * Handle %s routes
 * @param req - HTTP request
 * @param res - HTTP response
 * @returns Promise<boolean> - true if route was handled, false otherwise
 */
export async function handle%sRoutes(req: http.IncomingMessage, res: http.ServerResponse): Promise<boolean> {
    if (!req.url) return false;

    const parsedUrl = url.parse(req.url, true);
    const pathname = parsedUrl.pathname;
    const method = req.method?.toUpperCase();

    try {
        // GET /api/v1/admin/%s
        if (pathname === "/api/v1/admin/%s" && method === "GET") {
            await %sController.getAll(req, res);
            return true;
        }
        return false;
    } catch (error) {
        res.statusCode = 500;
        res.setHeader("Content-Type", "application/json");
        res.end(JSON.stringify({
            message: error instanceof Error ? error.message : "Internal server error"
        }));
        return true;
    }
}`, lowerName,
		upperName, lowerName, lowerName, lowerName)
}

func NodeController(lowerName string, upperName string) string {
	return fmt.Sprintf(`/**
   * Get the list of all %ss
   * @param req - HTTP request object
   * @param res - HTTP response object
   * @returns Promise<void> - Sends response directly to client
   * @throws {500} When an internal server error occurs during retrieval
   */
  public async getAll(req: http.IncomingMessage, res: http.ServerResponse): Promise<void> {
    try {
      const %ss = await this.%sService.findAll();

      // Set response headers
      res.setHeader("Content-Type", "application/json");

      if (!%ss || %ss.length === 0) {
        res.statusCode = 200;
        res.end(JSON.stringify({ 
          success: true, 
          data: [] 
        }));
        return;
      }

      res.statusCode = 200;
      res.end(JSON.stringify({ 
        success: true, 
        data: %ss 
      }));
    } catch (error) {
      res.statusCode = 500;
      res.setHeader("Content-Type", "application/json");
      res.end(JSON.stringify({
        success: false,
        message: error instanceof Error ? error.message : "Internal server error"
      }));
    }
  }`, lowerName,
		lowerName, lowerName, lowerName, lowerName, lowerName)
}

func NodeReqResMock() string {
	return `let mockRequest: Partial<http.IncomingMessage>;
  let mockResponse: Partial<http.ServerResponse>;`
}

func NodeReqResMockSetup() string {

	return `
  // Setup mock request
  mockRequest = {
    url: '/producs',
    method: 'GET',
    headers: {},
    };
    
    // Setup mock response
    mockResponse = {
      statusCode: 0,
      setHeader: jest.fn(),
      end: jest.fn(),
      };`
}

func NodeActController() string {
	return `await controller.getAll(mockRequest as http.IncomingMessage, mockResponse as http.ServerResponse);
`
}

func NodeAssertController(code int) string {
	return fmt.Sprintf(`expect(mockResponse.statusCode).toBe(%d);`, code)
}
