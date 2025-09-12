package exports

import "fmt"

// Return Express import
func ExpressImportation() string {
	return `
// Express importation
import express, {Request, Response} from "express";`
}

// Return Express export
func ExpressExportation() string {
	return `export default router;`
}

// Express router
func ExpressRouter() string {
	return `// Initialize Express router
const router = express.Router();`
}

// Return on exemple Express route endpoint
func ExpressRouteExemple(lowerName string) string {
	return fmt.Sprintf(`// Express Routes
// GET /api/v1/admin/%ss
router.get("/api/v1/admin/%ss", async (req: Request, res: Response) => { 
    %sController.getAll(req, res)
});
`, lowerName, lowerName, lowerName)
}

func ExpressController(lowerName string) string {
	return fmt.Sprintf(`/**
   * Get the list of all %ss
   * @param req 
   * @param res 
   * * @returns Promise<Response> - Express response object with appropriate status code and data
   * @throws {500} When an internal server error occurs during retrieval
   */
  public async getAll(req: Request, res: Response): Promise<Response> {
    try {
      const %ss = await this.%sService.findAll();
  
      if (!%ss || %ss.length === 0) return res.status(200).json({ success: true, data: [] });
  
      return res.status(200).json({ success: true, data: %ss });
    } catch (error) {
      return res.status(500).json(error instanceof Error ? error.message : "Internal server error");
    }
  }`, lowerName, lowerName, lowerName, lowerName, lowerName, lowerName)
}

func ExpressReqResMock() string {
	return `let mockRequest: Partial<Request>;
  let mockResponse: Partial<Response>;`
}

func ExpressReqResMockSetup() string {
	return `// Setup mock request
    mockRequest = {
      params: {},
      body: {},
      query: {},
      headers: {},
    };
    // Setup mock response with chainable methods
    mockResponse = {
      status: jest.fn().mockReturnThis(),
      json: jest.fn().mockReturnThis(),
      send: jest.fn().mockReturnThis(),
    };`
}

func ExpressActController() string {
	return ` await controller.getAll(mockRequest as Request, mockResponse as Response);`
}

func ExpressAssertController(code int) string {
	return fmt.Sprintf(`expect(mockResponse.status).toHaveBeenCalledWith(%d);`, code)
}
